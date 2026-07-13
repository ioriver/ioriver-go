package ioriver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const trafficData = `{
	"serviceStats": [
			{
					"serviceID": "%s",
					"points": []
			}
	],
	"granularity": "%s",
	"error": null
}`

func TestGetTraffic(t *testing.T) {
	path := fmt.Sprintf("/v2/%s%s/", trafficBasePath, testServiceId)
	serverData := fmt.Sprintf(trafficData, testServiceId, Day)
	expected := Traffic{ServiceStats: []ServiceStats{{ServiceId: testServiceId, Points: []Point{}}}, Granularity: "DAY", Error: ""}
	runGetTraffic(t, (*IORiverClient).GetTraffic, path, testServiceId, serverData, expected)
}

func TestRetrieveTrafficOvertime(t *testing.T) {
	path := fmt.Sprintf("/v2/%s", trafficBasePath)
	serverData := fmt.Sprintf(trafficData, testServiceId, Hour)
	expected := Traffic{ServiceStats: []ServiceStats{{ServiceId: testServiceId, Points: []Point{}}}, Granularity: "HOUR", Error: ""}

	setup()
	defer teardown()

	start := time.Now().UnixMilli()
	end := time.Now().UnixMilli()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected 'POST', got %s", r.Method)
		assert.Equal(t, fmt.Sprintf("%d", start), r.URL.Query().Get("startTime"))
		assert.Equal(t, fmt.Sprintf("%d", end), r.URL.Query().Get("endTime"))
		assert.Equal(t, Hour.String(), r.URL.Query().Get("granularity"))

		var requestPayload TrafficOvertimeRequest
		err := json.NewDecoder(r.Body).Decode(&requestPayload)
		assert.NoError(t, err)
		assert.Equal(t, []string{testServiceId}, requestPayload.ServiceIds)
		assert.Equal(t, []string{StatusCode.String(), HttpMethod.String(), HttpVersion.String()}, requestPayload.AdvancedMetricNames)

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	actual, err := client.RetrieveTrafficOvertime(
		[]string{testServiceId},
		start,
		end,
		Hour,
		[]AdvancedMetric{StatusCode, HttpMethod, HttpVersion},
	)
	if assert.NoError(t, err) {
		assert.Equal(t, &expected, actual)
	}
}

func TestGetFilteredMetrics(t *testing.T) {
	const providerName = "provider"
	const bytes = 100
	traffic := Traffic{ServiceStats: []ServiceStats{{ServiceId: testServiceId, Points: []Point{{Metrics: []Metric{{ProviderName: providerName, Metrics: Metrics{Bytes: bytes}}}}}}}, Granularity: "DAY", Error: ""}
	metrics := traffic.GetFilteredMetrics(testServiceId, func(metric *Metric, metricTimestamp int64) bool {
		return metric.ProviderName == providerName
	})
	if len(metrics) != 1 && metrics[0].Metrics.Bytes != bytes {
		t.Error("expected metric is not found")
	}
}

func runGetTraffic(
	t *testing.T,
	getter func(client *IORiverClient, serviceId string, startTime int64, endTime int64, granularity Granularity) (*Traffic, error),
	path string,
	serviceId string,
	serverData string,
	expected Traffic,
) {
	t.Helper()
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `%s`, serverData)
	}

	mux.HandleFunc(path, handler)

	start := time.Now().UnixMilli()
	end := time.Now().UnixMilli()

	actual, err := getter(client, serviceId, start, end, Day)
	if assert.NoError(t, err) {
		assert.Equal(t, &expected, actual)
	}
}
