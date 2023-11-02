package ioriver

import (
	"fmt"
	"testing"
)

const (
	testPerfMonitorName = "my monitor"
)

const serverPerfMonitorData = `{
	"id":"%s",
	"service":"%s",
	"name":"%s",
	"enabled":true,
	"url":"%s"
}`

var expectedPerfMonitor = PerformanceMonitor{
	Id:      testObjectId,
	Service: testServiceId,
	Name:    testPerfMonitorName,
	Url:     testUrl,
	Enabled: true,
}

func TestListPerfMonitors(t *testing.T) {
	path := fmt.Sprintf("/services/%s/performance-checks/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverPerfMonitorData, testObjectId, testServiceId,
		testPerfMonitorName, testUrl))
	expected := []PerformanceMonitor{expectedPerfMonitor}
	RunServiceList[PerformanceMonitor](t, (*IORiverClient).ListPerformanceMonitors, path, testServiceId, serverData, expected)
}

func TestGetPerfMonitor(t *testing.T) {
	path := fmt.Sprintf("/services/%s/performance-checks/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverPerfMonitorData, testObjectId, testServiceId,
		testPerfMonitorName, testUrl)
	RunServiceGet[PerformanceMonitor](t, (*IORiverClient).GetPerformanceMonitor, path, testServiceId, testObjectId, serverData,
		&expectedPerfMonitor)
}

func TestCreatePerfMonitor(t *testing.T) {
	newPerfMonitor := PerformanceMonitor{
		Service: testServiceId,
		Name:    testPerfMonitorName,
		Url:     testUrl,
	}

	path := fmt.Sprintf("/services/%s/performance-checks/", testServiceId)
	serverData := fmt.Sprintf(serverPerfMonitorData, testObjectId, testServiceId,
		testPerfMonitorName, testUrl)
	RunCreate[PerformanceMonitor](t, (*IORiverClient).CreatePerformanceMonitor, path, newPerfMonitor, serverData,
		&expectedPerfMonitor)
}
