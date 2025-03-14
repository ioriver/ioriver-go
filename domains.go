package ioriver

import (
	"fmt"
)

type DomainMappings struct {
	PathPattern string `json:"path_pattern,omitempty"`
	TargetId    string `json:"target_id"`
	TargetType  string `json:"target_type"`
}

type Domain struct {
	Id       string           `json:"id,omitempty"`
	Service  string           `json:"service"`
	Domain   string           `json:"domain"`
	Aliases  []string         `json:"aliases"`
	Mappings []DomainMappings `json:"mappings"`
}

const domainsBasePath = `services/%s/domains/`

func (client *IORiverClient) GetDomain(serviceId string, id string) (*Domain, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(domainsBasePath, serviceId), id)
	return Get[Domain](client, path)
}

func (client *IORiverClient) ListDomains(serviceId string) ([]Domain, error) {
	path := fmt.Sprintf(domainsBasePath, serviceId)
	return List[Domain](client, path)
}

func (client *IORiverClient) CreateDomain(domain Domain) (*Domain, error) {
	path := fmt.Sprintf(domainsBasePath, domain.Service)
	return Create[Domain](client, path, domain)
}

func (client *IORiverClient) UpdateDomain(domain Domain) (*Domain, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(domainsBasePath, domain.Service), domain.Id)
	return Update[Domain](client, path, domain)
}

func (client *IORiverClient) DeleteDomain(serviceId string, domainId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(domainsBasePath, serviceId), domainId)
	return Delete(client, path)
}
