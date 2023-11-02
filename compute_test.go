package ioriver

import (
	"fmt"
	"testing"
)

const (
	testComputeId       = "0aeea4ad-2498-44e8-85c3-c92b2ba858a2"
	testComputeName     = "test_compute_name"
	testComputeReqCode  = "async function onRequest(request) { console.log('test req'); }"
	testComputeRespCode = "async function onResponse(request, response){ console.log('test resp'); }"
	testRouteId         = "0aeea4ad-2498-44e8-85c3-c92b2ba858a3"
	testRouteDomain     = "test.ioriver-qa.com"
	testRoutePath       = "/api/*"
)

const serverComputeData = `{
	"id":"%s",
	"service":"%s",
	"name":"%s",
	"request_code":"%s",
	"response_code":"%s",
	"compute_routes": [
		{
			"id":"%s",
			"domain":"%s",
			"path":"%s"
		}
	]
}`

var expectedCompute = Compute{
	Id:           testComputeId,
	Service:      testServiceId,
	Name:         testComputeName,
	RequestCode:  testComputeReqCode,
	ResponseCode: testComputeRespCode,
	Routes: []ComputeRoute{
		{
			Id:     testRouteId,
			Domain: testRouteDomain,
			Path:   testRoutePath,
		},
	},
}

func TestListComputes(t *testing.T) {
	path := fmt.Sprintf("/services/%s/compute/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverComputeData, testComputeId, testServiceId,
		testComputeName, testComputeReqCode, testComputeRespCode, testRouteId, testRouteDomain, testRoutePath))
	expected := []Compute{expectedCompute}
	RunServiceList[Compute](t, (*IORiverClient).ListComputes, path,
		testServiceId, serverData, expected)
}

func TestGetCompute(t *testing.T) {
	path := fmt.Sprintf("/services/%s/compute/%s/", testServiceId, testComputeId)
	serverData := fmt.Sprintf(serverComputeData, testComputeId, testServiceId, testComputeName,
		testComputeReqCode, testComputeRespCode, testRouteId, testRouteDomain, testRoutePath)
	RunServiceGet[Compute](t, (*IORiverClient).GetCompute, path,
		testServiceId, testComputeId, serverData, &expectedCompute)
}

func TestCreateCompute(t *testing.T) {
	newCompute := Compute{
		Name:        testComputeName,
		Service:     testServiceId,
		RequestCode: "async function onRequest(request) { console.log('test req'); }",
		Routes: []ComputeRoute{
			{
				Domain: testRouteDomain,
				Path:   testRoutePath,
			},
		},
	}

	path := fmt.Sprintf("/services/%s/compute/", testServiceId)
	serverData := fmt.Sprintf(serverComputeData, testComputeId, testServiceId,
		testComputeName, testComputeReqCode, testComputeRespCode, testRouteId, testRouteDomain, testRoutePath)
	RunCreate[Compute](t, (*IORiverClient).CreateCompute, path, newCompute, serverData, &expectedCompute)
}
