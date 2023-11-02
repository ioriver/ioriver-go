package ioriver

import (
	"fmt"
	"testing"
)

const (
	testServiceProviderId = "c4fa41ca-6338-4035-a2a9-04928c510cc4"
	testWeight            = 52
	testStatus            = "Deploying"
	testStatusDetails     = "some service status"
)

const serverProviderData = `{
	"id":"%s",
	"account_provider":"%s",
	"service":"%s",
	"weight":%d,
	"is_failed": false,
	"status":"%s",
	"status_details":"%s",
	"restored": false,
	"cname": "example.com",
	"name": "Fastly",
	"display_name": "example",
	"is_unmanaged": true
}
`

var expectedServiceProvider = ServiceProvider{
	Id:              testServiceProviderId,
	AccountProvider: testAccountProviderId,
	Service:         testServiceId,
	Weight:          testWeight,
	IsFailed:        false,
	Status:          testStatus,
	StatusDetails:   testStatusDetails,
	Restored:        false,
	CName:           "example.com",
	Name:            "Fastly",
	DisplayName:     "example",
	IsUnmanaged:     true,
}

func TestListServiceProviders(t *testing.T) {
	path := fmt.Sprintf("/services/%s/providers/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverProviderData, testServiceProviderId,
		testAccountProviderId, testServiceId, testWeight, testStatus, testStatusDetails))
	expected := []ServiceProvider{expectedServiceProvider}
	RunServiceList[ServiceProvider](t, (*IORiverClient).ListServiceProviders, path,
		testServiceId, serverData, expected)
}

func TestGetServiceProvider(t *testing.T) {
	path := fmt.Sprintf("/services/%s/providers/%s/", testServiceId, testServiceProviderId)
	serverData := fmt.Sprintf(serverProviderData, testServiceProviderId, testAccountProviderId, testServiceId,
		testWeight, testStatus, testStatusDetails)
	RunServiceGet[ServiceProvider](t, (*IORiverClient).GetServiceProvider, path,
		testServiceId, testServiceProviderId, serverData, &expectedServiceProvider)
}

func TestCreateServiceProvider(t *testing.T) {
	newServiceProvider := ServiceProvider{
		AccountProvider: testAccountProviderId,
		Service:         testServiceId,
		Weight:          testWeight,
	}

	path := fmt.Sprintf("/services/%s/providers/", testServiceId)
	serverData := fmt.Sprintf(serverProviderData, testServiceProviderId, testAccountProviderId, testServiceId,
		testWeight, testStatus, testStatusDetails)
	RunCreate[ServiceProvider](t, (*IORiverClient).CreateServiceProvider, path,
		newServiceProvider, serverData, &expectedServiceProvider)
}
