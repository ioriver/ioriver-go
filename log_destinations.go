package ioriver

import (
	"fmt"
)

type DestinationType string

const (
	AWS_S3        DestinationType = "S3"
	S3_COMPATIBLE DestinationType = "S3_COMPATIBLE"
	HYDROLIX      DestinationType = "HYDROLIX"
)

type LogDestination struct {
	Id          string          `json:"id,omitempty"`
	Service     string          `json:"service"`
	Credentials interface{}     `json:"credentials,omitempty"`
	Name        string          `json:"name"`
	Type        DestinationType `json:"type"`
	S3Bucket    string          `json:"s3_bucket,omitempty"`
	S3Domain    string          `json:"s3_domain,omitempty"`
	S3Path      string          `json:"s3_path,omitempty"`
	S3Region    string          `json:"s3_region,omitempty"`
}

const logDestinationBasePath = `services/%s/log-destinations/`

func (client *IORiverClient) GetLogDestination(serviceId string, id string) (*LogDestination, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(logDestinationBasePath, serviceId), id)
	return Get[LogDestination](client, path)
}

func (client *IORiverClient) ListLogDestinations(serviceId string) ([]LogDestination, error) {
	path := fmt.Sprintf(logDestinationBasePath, serviceId)
	return List[LogDestination](client, path)
}

func (client *IORiverClient) CreateLogDestination(origin LogDestination) (*LogDestination, error) {
	path := fmt.Sprintf(logDestinationBasePath, origin.Service)
	return Create[LogDestination](client, path, origin)
}

func (client *IORiverClient) UpdateLogDestination(origin LogDestination) (*LogDestination, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(logDestinationBasePath, origin.Service), origin.Id)
	return Update[LogDestination](client, path, origin)
}

func (client *IORiverClient) DeleteLogDestination(serviceId string, originId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(logDestinationBasePath, serviceId), originId)
	return Delete(client, path)
}
