package ioriver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type IORiverClient struct {
	Token       string
	UserAgent   string
	EndpointUrl string
	Timeout     int
}

type CallParams struct {
	payload interface{}
	query   string
}

func NewClient(token string) *IORiverClient {
	return &IORiverClient{
		Token:       token,
		UserAgent:   userAgent,
		EndpointUrl: fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		Timeout:     180,
	}
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

func (client *IORiverClient) getAsyncTask(id int) (*AsyncTask, error) {

	url := client.EndpointUrl + "async_tasks/"
	httpClient := http.DefaultClient

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	if client.Token != "" {
		req.Header.Set("Authorization", "token "+client.Token)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		errDetails := ""
		if err == nil {
			errDetails = string(respBody)
		}

		respErr := fmt.Errorf("Request failed: %s, details: %s", resp.Status, errDetails)
		return nil, respErr
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var tasks []AsyncTask
	err = json.Unmarshal(respBody, &tasks)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling async-tasks json: %w", err)
	}

	for _, task := range tasks {
		if task.Id == id {
			return &task, nil
		}
	}

	return nil, nil
}

func (client *IORiverClient) waitForBackgrounTask(id int) error {

	TIMEOUT := 300
	elassped := 0
	var task *AsyncTask = nil
	var err error = nil

	for elassped < TIMEOUT {
		task, err = client.getAsyncTask(id)
		if err != nil {
			fmt.Printf("Error getting async-tasks: %s\n", err)
		}

		if err == nil && task != nil {
			if task.Status == "Status.COMPLETED" || task.Status == "Status.ERROR" {
				break
			}
		}

		time.Sleep(1 * time.Second)
		elassped += 1
	}

	if task != nil && task.Status == "Status.ERROR" {
		err = fmt.Errorf("Request failed: %s, details: %s", task.Message, task.Details)
	}

	return err
}

func (client *IORiverClient) CallApi(path string, method string, params CallParams) (*http.Response, error) {

	url := client.EndpointUrl + path
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

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	if client.Token != "" {
		req.Header.Set("Authorization", "token "+client.Token)
	}

	if reqBody != nil {
		req.Header.Set("content-type", "application/json")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		errDetails := ""
		if err == nil {
			errDetails = string(respBody)
		}

		respErr := fmt.Errorf("Request failed: %s, details: %s", resp.Status, errDetails)
		return nil, respErr
	}

	backgroundTask := resp.Header.Get("x-background-task-id")
	if backgroundTask != "" {
		id, err := strconv.Atoi(backgroundTask)
		if err == nil {
			err = client.waitForBackgrounTask(id)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

func Create[T interface{}, NewT interface{}](client *IORiverClient, path string, obj NewT) (*T, error) {
	resp, err := client.CallApi(path, "POST", CallParams{payload: obj})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json: %w", err)
	}
	return &result, nil
}

func Update[T interface{}](client *IORiverClient, path string, obj T) (*T, error) {
	resp, err := client.CallApi(path, "PUT", CallParams{payload: obj})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json: %w", err)
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

func Get[T interface{}](client *IORiverClient, path string) (*T, error) {

	resp, err := client.CallApi(path, "GET", CallParams{payload: nil})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var result T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("Unmarshaling json failed: %w", err)
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
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	var result []T
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return []T{}, fmt.Errorf("Error unmarshaling json: %w", err)
	}
	return result, nil
}
