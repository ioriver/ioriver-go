package ioriver

import (
	"fmt"
	"testing"
)

const (
	testTrafficPolicyId                = "0aeea4ad-2498-44e8-85c3-c92b2ba858a2"
	testTrafficPolicyServiceProviderId = "0aeea4ad-2498-44e8-85c3-c92b2ba858a4"
	testTrafficPolicyReqCode           = "async function onRequest(request) { console.log('test req'); }"
	testTrafficPolicyRespCode          = "async function onResponse(request, response){ console.log('test resp'); }"
)

const serverTrafficPolicyData = `{
	"id":"%s",
	"service":"%s",
    "type":"Static",
    "failover":false,
    "is_default":true,
    "providers":[
        {
            "service_provider":"%s",
            "weight":100    
        }
    ],
    "geos":[
        {
            "continent":null,
            "country":null,
            "subdivision":null
        }
    ],
    "health_checks":[],
    "performance_checks":[]
}`

var weight = 100

var expectedTrafficPolicy = TrafficPolicy{
	Id:        testTrafficPolicyId,
	Service:   testServiceId,
	Type:      "Static",
	Failover:  false,
	IsDefault: true,
	Providers: []TrafficPolicyProvider{
		{
			ServiceProvider: testTrafficPolicyServiceProviderId,
			Weight:          &weight,
		},
	},
	Geos: []TrafficPolicyGeo{
		{},
	},
	HealthChecks: []TrafficPolicyHealthCheck{},
	PerfChecks:   []TrafficPolicyPerfCheck{},
}

func TestListTrafficPolicies(t *testing.T) {
	path := fmt.Sprintf("/services/%s/traffic-policies/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverTrafficPolicyData, testTrafficPolicyId, testServiceId,
		testTrafficPolicyServiceProviderId))
	expected := []TrafficPolicy{expectedTrafficPolicy}
	RunServiceList[TrafficPolicy](t, (*IORiverClient).ListTrafficPolicies, path,
		testServiceId, serverData, expected)
}

func TestGetTrafficPolicy(t *testing.T) {
	path := fmt.Sprintf("/services/%s/traffic-policies/%s/", testServiceId, testTrafficPolicyId)
	serverData := fmt.Sprintf(serverTrafficPolicyData, testTrafficPolicyId, testServiceId, testTrafficPolicyServiceProviderId)
	RunServiceGet[TrafficPolicy](t, (*IORiverClient).GetTrafficPolicy, path,
		testServiceId, testTrafficPolicyId, serverData, &expectedTrafficPolicy)
}

func TestCreateTrafficPolicy(t *testing.T) {
	var newTrafficPolicy = TrafficPolicy{
		Id:        testTrafficPolicyId,
		Service:   testServiceId,
		Type:      "Static",
		Failover:  false,
		IsDefault: true,
		Providers: []TrafficPolicyProvider{
			{
				ServiceProvider: testTrafficPolicyServiceProviderId,
				Weight:          &weight,
			},
		},
		Geos: []TrafficPolicyGeo{
			{},
		},
		HealthChecks: []TrafficPolicyHealthCheck{},
		PerfChecks:   []TrafficPolicyPerfCheck{},
	}

	path := fmt.Sprintf("/services/%s/traffic-policies/", testServiceId)
	serverData := fmt.Sprintf(serverTrafficPolicyData, testTrafficPolicyId, testServiceId, testTrafficPolicyServiceProviderId)
	RunCreate[TrafficPolicy](t, (*IORiverClient).CreateTrafficPolicy, path,
		newTrafficPolicy, serverData, &expectedTrafficPolicy)
}
