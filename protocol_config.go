package ioriver

import "fmt"

type ProtocolConfig struct {
	Id           string `json:"id,omitempty"`
	Service      string `json:"service"`
	Http2Enabled bool   `json:"http2_enabled"`
	Http3Enabled bool   `json:"http3_enabled"`
	Ipv6Enabled  bool   `json:"ipv6_enabled"`
}

const protocolConfigBasePath = `services/%s/protocol-config/`

func (client *IORiverClient) GetProtocolConfig(serviceId string, id string) (*ProtocolConfig, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(protocolConfigBasePath, serviceId), id)
	return Get[ProtocolConfig](client, path)
}

func (client *IORiverClient) ListProtocolConfigs(serviceId string) ([]ProtocolConfig, error) {
	path := fmt.Sprintf(protocolConfigBasePath, serviceId)
	return List[ProtocolConfig](client, path)
}

func (client *IORiverClient) CreateProtocolConfig(protocolConfig ProtocolConfig) (*ProtocolConfig, error) {
	path := fmt.Sprintf(protocolConfigBasePath, protocolConfig.Service)
	return Create[ProtocolConfig](client, path, protocolConfig)
}

func (client *IORiverClient) UpdateProtocolConfig(protocolConfig ProtocolConfig) (*ProtocolConfig, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(protocolConfigBasePath, protocolConfig.Service), protocolConfig.Id)
	return Update[ProtocolConfig](client, path, protocolConfig)
}

func (client *IORiverClient) DeleteProtocolConfig(serviceId string, protocolConfigId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(protocolConfigBasePath, serviceId), protocolConfigId)
	return Delete(client, path)
}
