package ioriver

import (
	"fmt"
	"testing"
)

const (
	testDomainId = "0aeea4ad-2498-44e8-85c3-c92b2ba858a2"
)

const serverDomainData = `{
	"id":"%s",
	"service":"%s",
	"domain":"%s",
	"path_pattern":"%s",
	"origin":"%s",
	"load_balancer": ""
}`

var expectedDomain = Domain{
	Id:           testDomainId,
	Service:      testServiceId,
	Domain:       testDomainName,
	PathPattern:  testPathPattern,
	Origin:       testObjectId,
	LoadBalancer: "",
}

func TestListDomains(t *testing.T) {
	path := fmt.Sprintf("/services/%s/domains/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverDomainData, testDomainId, testServiceId, testDomainName,
		testPathPattern, testObjectId))
	expected := []Domain{expectedDomain}
	RunServiceList[Domain](t, (*IORiverClient).ListDomains, path, testServiceId, serverData, expected)
}

func TestGetDomain(t *testing.T) {
	path := fmt.Sprintf("/services/%s/domains/%s/", testServiceId, testDomainId)
	serverData := fmt.Sprintf(serverDomainData, testDomainId, testServiceId, testDomainName,
		testPathPattern, testObjectId)
	RunServiceGet[Domain](t, (*IORiverClient).GetDomain, path, testServiceId, testDomainId, serverData, &expectedDomain)
}

func TestCreateDomain(t *testing.T) {
	newDomain := Domain{
		Domain:      testDomainName,
		Service:     testServiceId,
		Origin:      testObjectId,
		PathPattern: testPathPattern,
	}

	path := fmt.Sprintf("/services/%s/domains/", testServiceId)
	serverData := fmt.Sprintf(serverDomainData, testDomainId, testServiceId, testDomainName,
		testPathPattern, testObjectId)
	RunCreate[Domain](t, (*IORiverClient).CreateDomain, path, newDomain, serverData, &expectedDomain)
}
