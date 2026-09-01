package ioriver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type IORiverClient struct {
	Token            string
	UserAgent        string
	EndpointUrl      string
	TerraformVersion string
	Timeout          int
}

type CallParams struct {
	payload interface{}
	query   string
}

type APIError struct {
	StatusCode int
	Status     string
	RequestID  string
	Details    string
}

func (e *APIError) Error() string {
	requestID := e.RequestID
	if requestID == "" {
		requestID = "unknown"
	}

	return fmt.Sprintf("request failed: %s, request-id: %s, details: %s", e.Status, requestID, e.Details)
}

func NewClient(token string) *IORiverClient {
	c := &IORiverClient{
		Token:       token,
		UserAgent:   userAgent,
		EndpointUrl: fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		Timeout:     180,
	}
	return c
}

type AsyncTask struct {
	Id       int    `json:"id"`
	Status   string `json:"status"`
	Progress int    `json:"progress"`
	Message  string `json:"message,omitempty"`
	Details  string `json:"details,omitempty"`
	Created  string `json:"created,omitempty"`
	Title    string `json:"title,omitempty"`
}

func normalizeAsyncTaskStatus(status string) string {
	return strings.TrimPrefix(status, "Status.")
}

func (client *IORiverClient) getAsyncTask(id int) (*AsyncTask, error) {

	url := client.EndpointUrl + "v1/async_task_by_id/" + strconv.Itoa(id) + "/"
	httpClient := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	if client.Token != "" {
		req.Header.Set("Authorization", "token "+client.Token)
	}

	if client.TerraformVersion != "" {
		req.Header.Set("X-Terraform-Version", client.TerraformVersion)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Not found task treated as success
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(resp.Body)
		errDetails := ""
		if err == nil {
			errDetails = string(respBody)
		}

		requestId := resp.Header.Get("x-request-id")
		respErr := &APIError{StatusCode: resp.StatusCode, Status: resp.Status, RequestID: requestId, Details: errDetails}
		return nil, respErr
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var task AsyncTask
	err = json.Unmarshal(respBody, &task)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling async-tasks json: %w", err)
	}

	if task.Id == id {
		return &task, nil
	}
	return nil, nil
}

func (client *IORiverClient) waitForBackgroundTask(id int, requestId string) error {
	waits := []int{1, 2, 4, 8, 10}
	pollIndex := 0
	elapsed := 0
	var task *AsyncTask = nil
	var err error = nil

	for elapsed < defaultAsyncTaskTimeout {
		task, err = client.getAsyncTask(id)
		if err != nil {
			break
		}

		if task == nil {
			fmt.Fprintf(os.Stderr, "Async task with id %d not found. \n", id)
			err = fmt.Errorf("async task not found: id=%d, request-id=%s", id, requestId)
			break
		}

		status := normalizeAsyncTaskStatus(task.Status)
		if status == "COMPLETED" || status == "ERROR" {
			break
		}

		sleepFor := waits[pollIndex]
		time.Sleep(time.Duration(sleepFor) * time.Second)
		elapsed += sleepFor

		if pollIndex < len(waits)-1 {
			pollIndex += 1
		}
	}

	if task != nil {
		switch normalizeAsyncTaskStatus(task.Status) {
		case "ERROR":
			err = &APIError{StatusCode: 0, Status: task.Message, RequestID: requestId, Details: task.Details}
		case "COMPLETED":
			err = nil
		default:
			err = fmt.Errorf("request did not complete within timeout, current status: %s, request-id: %s", task.Status, requestId)
		}
	}

	return err
}

func (client *IORiverClient) CallApi(path string, method string, params CallParams) (*http.Response, error) {

	version := getVersion(path)
	url := fmt.Sprintf("%s%s/%s", client.EndpointUrl, version, path)
	httpClient := http.DefaultClient

	var reqBody io.Reader = nil
	if params.payload != nil {
		var jsonBody []byte
		jsonBody, err := json.Marshal(params.payload)
		if err != nil {
			return nil, fmt.Errorf("error marshalling payload to JSON: %w", err)
		}

		reqBody = bytes.NewReader(jsonBody)
	}

	if params.query != "" {
		url += "?" + params.query
	}

	fmt.Fprintf(os.Stderr, "[IORIVER API CALL REQ] %s %s\n", method, url)

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	if client.Token != "" {
		req.Header.Set("Authorization", "token "+client.Token)
	}

	if client.TerraformVersion != "" {
		req.Header.Set("X-Terraform-Version", client.TerraformVersion)
	}

	if reqBody != nil {
		req.Header.Set("content-type", "application/json")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IORIVER API CALL] %s %s -> ERROR: %s\n", method, url, err)
		return resp, err
	}

	requestId := resp.Header.Get("x-request-id")
	if requestId == "" {
		requestId = "unknown"
	}
	fmt.Fprintf(os.Stderr, "[IORIVER API CALL RESP] %s %s -> %s (req-id: %s)\n", method, url, resp.Status, requestId)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		errDetails := ""
		if err == nil {
			errDetails = string(respBody)
		}

		respErr := &APIError{StatusCode: resp.StatusCode, Status: resp.Status, RequestID: requestId, Details: errDetails}
		return nil, respErr
	}

	backgroundTask := resp.Header.Get("x-background-task-id")
	if backgroundTask != "" {
		id, err := strconv.Atoi(backgroundTask)
		if err == nil {
			err = client.waitForBackgroundTask(id, requestId)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

func Create[T interface{}, NewT interface{}](client *IORiverClient, path string, obj NewT, queryParams ...string) (*T, error) {
	resp, err := client.CallApi(path, "POST", CallParams{payload: obj, query: strings.Join(queryParams, "&")})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return &result, nil
}

func Update[T interface{}, NewT interface{}](client *IORiverClient, path string, obj NewT) (*T, error) {
	resp, err := client.CallApi(path, "PUT", CallParams{payload: obj})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return &result, nil
}

func Delete(client *IORiverClient, path string) error {
	_, err := client.CallApi(path, "DELETE", CallParams{payload: nil})
	return err
}

func DeleteWithQueryString(client *IORiverClient, path string, query string) error {
	_, err := client.CallApi(path, "DELETE", CallParams{payload: nil, query: query})
	return err
}

func Get[T interface{}](client *IORiverClient, path string, queryParams ...string) (*T, error) {

	callParams := CallParams{payload: nil, query: strings.Join(queryParams, "&")}
	resp, err := client.CallApi(path, "GET", callParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return &result, nil
}

func List[T interface{}](client *IORiverClient, path string) ([]T, error) {

	resp, err := client.CallApi(path, "GET", CallParams{payload: nil})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result []T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return []T{}, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return result, nil
}

func getVersion(path string) string {
	if strings.HasPrefix(path, trafficBasePath) {
		return "v2"
	}
	return "v1"
}
