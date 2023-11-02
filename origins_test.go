package ioriver

import (
	"fmt"
	"testing"
)

const (
	testOriginHost = "www.example.com"
	testOriginPath = "/test/path"
)

const serverOriginData = `{
	"id":"%s",
	"service":"%s",
	"host":"%s",
	"port":443,
	"protocol":"HTTPS",
	"path": "%s"
}`

var expectedOrigin = Origin{
	Id:       testObjectId,
	Service:  testServiceId,
	Host:     testOriginHost,
	Port:     443,
	Protocol: "HTTPS",
	Path:     testOriginPath,
}

func TestListOrigins(t *testing.T) {
	path := fmt.Sprintf("/services/%s/origins/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverOriginData, testObjectId, testServiceId, testOriginHost, testOriginPath))
	expected := []Origin{expectedOrigin}
	RunServiceList[Origin](t, (*IORiverClient).ListOrigins, path, testServiceId, serverData, expected)
}

func TestGetOrigin(t *testing.T) {
	path := fmt.Sprintf("/services/%s/origins/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverOriginData, testObjectId, testServiceId, testOriginHost, testOriginPath)
	RunServiceGet[Origin](t, (*IORiverClient).GetOrigin, path, testServiceId, testObjectId, serverData, &expectedOrigin)
}

func TestCreateOrigin(t *testing.T) {
	newOrigin := Origin{
		Host:     testOriginHost,
		Service:  testServiceId,
		Port:     443,
		Protocol: "HTTPS",
		Path:     testOriginPath,
	}

	path := fmt.Sprintf("/services/%s/origins/", testServiceId)
	serverData := fmt.Sprintf(serverOriginData, testObjectId, testServiceId, testOriginHost, testOriginPath)
	RunCreate[Origin](t, (*IORiverClient).CreateOrigin, path, newOrigin, serverData, &expectedOrigin)
}
