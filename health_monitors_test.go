package ioriver

import (
	"fmt"
	"testing"
)

const (
	testHealthMonitorName = "my monitor"
)

const serverHealthMonitorData = `{
	"id":"%s",
	"service":"%s",
	"name":"%s",
	"enabled":true,
	"url":"%s"
}`

var expectedHealthMonitor = HealthMonitor{
	Id:      testObjectId,
	Service: testServiceId,
	Name:    testHealthMonitorName,
	Url:     testUrl,
	Enabled: true,
}

func TestListHealthMonitors(t *testing.T) {
	path := fmt.Sprintf("/services/%s/health-checks/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverHealthMonitorData, testObjectId, testServiceId,
		testHealthMonitorName, testUrl))
	expected := []HealthMonitor{expectedHealthMonitor}
	RunServiceList[HealthMonitor](t, (*IORiverClient).ListHealthMonitors, path, testServiceId, serverData, expected)
}

func TestGetHealthMonitor(t *testing.T) {
	path := fmt.Sprintf("/services/%s/health-checks/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverHealthMonitorData, testObjectId, testServiceId,
		testHealthMonitorName, testUrl)
	RunServiceGet[HealthMonitor](t, (*IORiverClient).GetHealthMonitor, path, testServiceId, testObjectId, serverData,
		&expectedHealthMonitor)
}

func TestCreateHealthMonitor(t *testing.T) {
	newHealthMonitor := HealthMonitor{
		Service: testServiceId,
		Name:    testHealthMonitorName,
		Url:     testUrl,
	}

	path := fmt.Sprintf("/services/%s/health-checks/", testServiceId)
	serverData := fmt.Sprintf(serverHealthMonitorData, testObjectId, testServiceId,
		testHealthMonitorName, testUrl)
	RunCreate[HealthMonitor](t, (*IORiverClient).CreateHealthMonitor, path, newHealthMonitor, serverData,
		&expectedHealthMonitor)
}
