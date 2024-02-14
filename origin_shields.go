package ioriver

import "fmt"

type ProviderOriginShield struct {
	ServiceProvider  string `json:"service_provider"`
	ProviderLocation string `json:"provider_location,omitempty"`
}

type OriginShieldLocation struct {
	Country     string `json:"country"`
	Subdivision string `json:"subdivision,omitempty"`
}

type OriginShield struct {
	Id        string                 `json:"id,omitempty"`
	Service   string                 `json:"service"`
	Location  OriginShieldLocation   `json:"location"`
	Providers []ProviderOriginShield `json:"providers"`
}

const originShieldBasePath = `services/%s/origin-shield/`

func (client *IORiverClient) GetOriginShield(serviceId string, id string) (*OriginShield, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originShieldBasePath, serviceId), id)
	return Get[OriginShield](client, path)
}

func (client *IORiverClient) ListOriginShields(serviceId string) ([]OriginShield, error) {
	path := fmt.Sprintf(originShieldBasePath, serviceId)
	return List[OriginShield](client, path)
}

func (client *IORiverClient) CreateOriginShield(policy OriginShield) (*OriginShield, error) {
	path := fmt.Sprintf(originShieldBasePath, policy.Service)
	return Create[OriginShield](client, path, policy)
}

func (client *IORiverClient) UpdateOriginShield(policy OriginShield) (*OriginShield, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originShieldBasePath, policy.Service), policy.Id)
	return Update[OriginShield](client, path, policy)
}

func (client *IORiverClient) DeleteOriginShield(serviceId string, policyId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originShieldBasePath, serviceId), policyId)
	return Delete(client, path)
}
