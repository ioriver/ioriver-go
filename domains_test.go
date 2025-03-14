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
	"mappings": [
		{
			"path_pattern":"%s",
			"target_id":"%s",
			"target_type":"origin"
    }
	]
}`

var expectedDomain = Domain{
	Id:      testDomainId,
	Service: testServiceId,
	Domain:  testDomainName,
	Mappings: []DomainMappings{
		{
			PathPattern: testPathPattern,
			TargetId:    testObjectId,
			TargetType:  "origin",
		},
	},
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
		Domain:  testDomainName,
		Service: testServiceId,
		Mappings: []DomainMappings{
			{
				PathPattern: testPathPattern,
				TargetId:    testObjectId,
				TargetType:  "origin",
			},
		},
	}

	path := fmt.Sprintf("/services/%s/domains/", testServiceId)
	serverData := fmt.Sprintf(serverDomainData, testDomainId, testServiceId, testDomainName,
		testPathPattern, testObjectId)
	RunCreate[Domain](t, (*IORiverClient).CreateDomain, path, newDomain, serverData, &expectedDomain)
}
