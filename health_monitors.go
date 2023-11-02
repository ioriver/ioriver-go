package ioriver

import (
	"fmt"
)

type HealthMonitor struct {
	Id      string `json:"id,omitempty"`
	Service string `json:"service"`
	Name    string `json:"name"`

	Enabled bool `json:"enabled,omitempty"`

	Url string `json:"url"`
}

const healthMonBasePath = `services/%s/health-checks/`

func (client *IORiverClient) GetHealthMonitor(serviceId string, id string) (*HealthMonitor, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(healthMonBasePath, serviceId), id)
	return Get[HealthMonitor](client, path)
}

func (client *IORiverClient) ListHealthMonitors(serviceId string) ([]HealthMonitor, error) {
	path := fmt.Sprintf(healthMonBasePath, serviceId)
	return List[HealthMonitor](client, path)
}

func (client *IORiverClient) CreateHealthMonitor(healthMonitor HealthMonitor) (*HealthMonitor, error) {
	path := fmt.Sprintf(healthMonBasePath, healthMonitor.Service)
	return Create[HealthMonitor](client, path, healthMonitor)
}

func (client *IORiverClient) UpdateHealthMonitor(healthMonitor HealthMonitor) (*HealthMonitor, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(healthMonBasePath, healthMonitor.Service), healthMonitor.Id)
	return Update[HealthMonitor](client, path, healthMonitor)
}

func (client *IORiverClient) DeleteHealthMonitor(serviceId string, healthMonitorId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(healthMonBasePath, serviceId), healthMonitorId)
	return Delete(client, path)
}
