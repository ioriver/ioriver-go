package ioriver

import (
	"fmt"
	"testing"
)

const (
	testOriginShieldId                = "0aeea4ad-2498-44e8-85c3-c92b2ba858a2"
	testOriginShieldServiceProviderId = "0aeea4ad-2498-44e8-85c3-c92b2ba858a4"
)

const serverOriginShieldData = `{
	"id":"%s",
	"service":"%s",
	"location":{"country":"DE","subdivision":null},
	"providers":[
			{
					"service_provider":"%s",
					"provider_location":"some-location" 
			}
	]
}`

var expectedOriginShield = OriginShield{
	Id:      testOriginShieldId,
	Service: testServiceId,
	Location: OriginShieldLocation{
		Country: "DE",
	},
	Providers: []ProviderOriginShield{
		{
			ServiceProvider:  testOriginShieldServiceProviderId,
			ProviderLocation: "some-location",
		},
	},
}

func TestListOriginShields(t *testing.T) {
	path := fmt.Sprintf("/services/%s/origin-shield/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverOriginShieldData, testOriginShieldId, testServiceId,
		testOriginShieldServiceProviderId))
	expected := []OriginShield{expectedOriginShield}
	RunServiceList[OriginShield](t, (*IORiverClient).ListOriginShields, path,
		testServiceId, serverData, expected)
}

func TestGetOriginShield(t *testing.T) {
	path := fmt.Sprintf("/services/%s/origin-shield/%s/", testServiceId, testOriginShieldId)
	serverData := fmt.Sprintf(serverOriginShieldData, testOriginShieldId, testServiceId, testOriginShieldServiceProviderId)
	RunServiceGet[OriginShield](t, (*IORiverClient).GetOriginShield, path,
		testServiceId, testOriginShieldId, serverData, &expectedOriginShield)
}

func TestCreateOriginShield(t *testing.T) {
	var newOriginShield = OriginShield{
		Id:      testOriginShieldId,
		Service: testServiceId,
		Providers: []ProviderOriginShield{
			{
				ServiceProvider: testOriginShieldServiceProviderId,
			},
		},
	}

	path := fmt.Sprintf("/services/%s/origin-shield/", testServiceId)
	serverData := fmt.Sprintf(serverOriginShieldData, testOriginShieldId, testServiceId, testOriginShieldServiceProviderId)
	RunCreate[OriginShield](t, (*IORiverClient).CreateOriginShield, path,
		newOriginShield, serverData, &expectedOriginShield)
}
