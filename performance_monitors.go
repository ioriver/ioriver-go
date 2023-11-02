package ioriver

import (
	"fmt"
)

type PerformanceMonitor struct {
	Id      string `json:"id,omitempty"`
	Service string `json:"service"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled,omitempty"`
	Url     string `json:"url"`
}

const perfMonBasePath = `services/%s/performance-checks/`

func (client *IORiverClient) GetPerformanceMonitor(serviceId string, id string) (*PerformanceMonitor, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(perfMonBasePath, serviceId), id)
	return Get[PerformanceMonitor](client, path)
}

func (client *IORiverClient) ListPerformanceMonitors(serviceId string) ([]PerformanceMonitor, error) {
	path := fmt.Sprintf(perfMonBasePath, serviceId)
	return List[PerformanceMonitor](client, path)
}

func (client *IORiverClient) CreatePerformanceMonitor(perfMonitor PerformanceMonitor) (*PerformanceMonitor, error) {
	path := fmt.Sprintf(perfMonBasePath, perfMonitor.Service)
	return Create[PerformanceMonitor](client, path, perfMonitor)
}

func (client *IORiverClient) UpdatePerformanceMonitor(perfMonitor PerformanceMonitor) (*PerformanceMonitor, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(perfMonBasePath, perfMonitor.Service), perfMonitor.Id)
	return Update[PerformanceMonitor](client, path, perfMonitor)
}

func (client *IORiverClient) DeletePerformanceMonitor(serviceId string, perfMonitorId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(perfMonBasePath, serviceId), perfMonitorId)
	return Delete(client, path)
}
