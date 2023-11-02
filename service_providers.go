package ioriver

import (
	"fmt"
	"time"
)

type ServiceProvider struct {
	Id              string `json:"id,omitempty"`
	AccountProvider string `json:"account_provider,omitempty"`
	Service         string `json:"service"`
	Weight          int    `json:"weight"`
	IsUnmanaged     bool   `json:"is_unmanaged,omitempty"`
	CName           string `json:"cname,omitempty"`
	DisplayName     string `json:"display_name,omitempty"`
	IsFailed        bool   `json:"is_failed,omitempty"`
	Status          string `json:"status,omitempty"`
	StatusDetails   string `json:"status_details,omitempty"`
	Restored        bool   `json:"restored,omitempty"`
	Name            string `json:"name,omitempty"`
}

const spBasePath = `services/%s/providers/`

func (client *IORiverClient) GetServiceProvider(serviceId string, id string) (*ServiceProvider, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(spBasePath, serviceId), id)
	return Get[ServiceProvider](client, path)
}

func (client *IORiverClient) ListServiceProviders(serviceId string) ([]ServiceProvider, error) {
	path := fmt.Sprintf(spBasePath, serviceId)
	return List[ServiceProvider](client, path)
}

func (client *IORiverClient) CreateServiceProvider(provider ServiceProvider) (*ServiceProvider, error) {
	path := fmt.Sprintf(spBasePath, provider.Service)
	return Create[ServiceProvider](client, path, provider)
}

func (client *IORiverClient) UpdateServiceProvider(provider ServiceProvider) (*ServiceProvider, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(spBasePath, provider.Service), provider.Id)
	return Update[ServiceProvider](client, path, provider)
}

func (client *IORiverClient) DeleteServiceProvider(serviceId string, providerId string, force string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(spBasePath, serviceId), providerId)
	queryString := "force=" + force
	return DeleteWithQueryString(client, path, queryString)
}

func (client *IORiverClient) WaitServiceProviderReady(serviceId string, id string) error {
	elaspedTime := 0
	timeoutDuration := 30 * time.Minute

	for {
		serviceProvider, err := client.GetServiceProvider(serviceId, id)
		if err != nil {
			return err
		}

		if serviceProvider.Status == "Active" {
			return nil
		}

		time.Sleep(10 * time.Second)
		elaspedTime += 10

		if elaspedTime >= int(timeoutDuration.Seconds()) {
			return fmt.Errorf("Timeout waiting for serivce-provider %s to become active", id)
		}
	}
}
