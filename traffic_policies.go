package ioriver

import (
	"fmt"
)

type TrafficPolicyProvider struct {
	ServiceProvider string `json:"service_provider"`
	Weight          int    `json:"weight"`
}

type TrafficPolicyGeo struct {
	Continent   string `json:"continent,omitempty"`
	Country     string `json:"country,omitempty"`
	Subdivision string `json:"subdivision,omitempty"`
}

type TrafficPolicyHealthCheck struct {
	HealthCheck string `json:"health_check"`
}

type TrafficPolicyPerfCheck struct {
	PerformanceCheck string `json:"performance_check"`
}

type TrafficPolicy struct {
	Id           string                     `json:"id,omitempty"`
	Service      string                     `json:"service"`
	Type         string                     `json:"type"`
	Failover     bool                       `json:"failover"`
	IsDefault    bool                       `json:"is_default"`
	Providers    []TrafficPolicyProvider    `json:"providers"`
	Geos         []TrafficPolicyGeo         `json:"geos"`
	HealthChecks []TrafficPolicyHealthCheck `json:"health_checks"`
	PerfChecks   []TrafficPolicyPerfCheck   `json:"performance_checks"`
}

const trafficPolicyBasePath = `services/%s/traffic-policies/`

func (client *IORiverClient) GetTrafficPolicy(serviceId string, id string) (*TrafficPolicy, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(trafficPolicyBasePath, serviceId), id)
	return Get[TrafficPolicy](client, path)
}

func (client *IORiverClient) ListTrafficPolicies(serviceId string) ([]TrafficPolicy, error) {
	path := fmt.Sprintf(trafficPolicyBasePath, serviceId)
	return List[TrafficPolicy](client, path)
}

func (client *IORiverClient) CreateTrafficPolicy(policy TrafficPolicy) (*TrafficPolicy, error) {
	path := fmt.Sprintf(trafficPolicyBasePath, policy.Service)
	return Create[TrafficPolicy](client, path, policy)
}

func (client *IORiverClient) UpdateTrafficPolicy(policy TrafficPolicy) (*TrafficPolicy, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(trafficPolicyBasePath, policy.Service), policy.Id)
	return Update[TrafficPolicy](client, path, policy)
}

func (client *IORiverClient) DeleteTrafficPolicy(serviceId string, policyId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(trafficPolicyBasePath, serviceId), policyId)
	return Delete(client, path)
}
