package ioriver

import (
	"fmt"
)

type ServiceConfig struct {
	Description   string                 `json:"description,omitempty"`
	Version       int                    `json:"version,omitempty"`
	ParentVersion int                    `json:"parent_version,omitempty"`
	ConfigJSON    map[string]interface{} `json:"config_json,omitempty"`
}

type ServiceCreateFromConfigRequest struct {
	Description   string                 `json:"description"`
	CertificateID string                 `json:"certificate_id"`
	ServiceConfig map[string]interface{} `json:"service_config"`
}

type ServiceCreateFromConfigResponse struct {
	ID            string   `json:"id"`
	Account       string   `json:"account,omitempty"`
	Name          string   `json:"name"`
	Version       int      `json:"version,omitempty"`
	Description   string   `json:"description,omitempty"`
	Certificate   string   `json:"certificate,omitempty"`
	Certificates  []string `json:"certificates,omitempty"`
	ServiceUid    string   `json:"service_uid,omitempty"`
	Cname         string   `json:"cname,omitempty"`
	ReadOnly      bool     `json:"read_only,omitempty"`
	ActiveVersion int      `json:"active_version,omitempty"`
}

func (client *IORiverClient) GetCurrentServiceConfig(serviceID string) (*ServiceConfig, error) {
	path := fmt.Sprintf("%s%s/current_service_config/", servicesBasePath, serviceID)
	return Get[ServiceConfig](client, path)
}

func (client *IORiverClient) UpdateServiceConfig(serviceID string, serviceConfig ServiceConfig) (*ServiceConfig, error) {
	pathUpdateConfig := fmt.Sprintf("%s%s/service-configs/", servicesBasePath, serviceID)
	return Create[ServiceConfig](client, pathUpdateConfig, serviceConfig)
}
