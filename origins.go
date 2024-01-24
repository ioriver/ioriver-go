package ioriver

import (
	"fmt"
)

type Origin struct {
	Id        string `json:"id,omitempty"`
	Service   string `json:"service"`
	Host      string `json:"host"`
	Protocol  string `json:"protocol,omitempty"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
	IsS3      bool   `json:"is_s3,omitempty"`
	TimeoutMs int    `json:"timeout_ms,omitempty"`
}

const originsBasePath = `services/%s/origins/`

func (client *IORiverClient) GetOrigin(serviceId string, id string) (*Origin, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originsBasePath, serviceId), id)
	return Get[Origin](client, path)
}

func (client *IORiverClient) ListOrigins(serviceId string) ([]Origin, error) {
	path := fmt.Sprintf(originsBasePath, serviceId)
	return List[Origin](client, path)
}

func (client *IORiverClient) CreateOrigin(origin Origin) (*Origin, error) {
	if origin.Protocol == "HTTPS" {
		origin.Port = 443
	} else {
		origin.Port = 80
	}
	path := fmt.Sprintf(originsBasePath, origin.Service)
	return Create[Origin](client, path, origin)
}

func (client *IORiverClient) UpdateOrigin(origin Origin) (*Origin, error) {
	if origin.Protocol == "HTTPS" {
		origin.Port = 443
	} else {
		origin.Port = 80
	}
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originsBasePath, origin.Service), origin.Id)
	return Update[Origin](client, path, origin)
}

func (client *IORiverClient) DeleteOrigin(serviceId string, originId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originsBasePath, serviceId), originId)
	return Delete(client, path)
}
