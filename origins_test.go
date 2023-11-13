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
	"protocol":"HTTPS",
	"path": "%s",
	"is_s3": true
}`

var expectedOrigin = Origin{
	Id:       testObjectId,
	Service:  testServiceId,
	Host:     testOriginHost,
	Protocol: "HTTPS",
	Path:     testOriginPath,
	IsS3:     true,
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
		Protocol: "HTTPS",
		Path:     testOriginPath,
		IsS3:     true,
	}

	path := fmt.Sprintf("/services/%s/origins/", testServiceId)
	serverData := fmt.Sprintf(serverOriginData, testObjectId, testServiceId, testOriginHost, testOriginPath)
	RunCreate[Origin](t, (*IORiverClient).CreateOrigin, path, newOrigin, serverData, &expectedOrigin)
}
