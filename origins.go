package ioriver

import (
	"fmt"
)

type OriginShieldProvider struct {
	ServiceProvider  string `json:"service_provider"`
	ProviderLocation string `json:"provider_location,omitempty"`
}

type OriginShieldLocation struct {
	Country     string `json:"country"`
	Subdivision string `json:"subdivision,omitempty"`
}

type Origin struct {
	Id              string                 `json:"id,omitempty"`
	Service         string                 `json:"service"`
	Host            string                 `json:"host"`
	Protocol        string                 `json:"protocol,omitempty"`
	Path            string                 `json:"path,omitempty"`
	HttpPort        int                    `json:"http_port,omitempty"`
	HttpsPort       int                    `json:"https_port,omitempty"`
	IsS3            bool                   `json:"is_s3,omitempty"`
	IsPrivateS3     bool                   `json:"is_private_s3,omitempty"`
	S3BucketName    string                 `json:"s3_bucket_name,omitempty"`
	S3AwsRegion     string                 `json:"s3_aws_region,omitempty"`
	S3AwsKey        string                 `json:"s3_aws_key,omitempty"`
	S3AwsSecret     string                 `json:"s3_aws_secret,omitempty"`
	TimeoutMs       int                    `json:"timeout_ms,omitempty"`
	VerifyTLS       bool                   `json:"verify_tls,omitempty"`
	ShieldLocation  *OriginShieldLocation  `json:"shield_location,omitempty"`
	ShieldProviders []OriginShieldProvider `json:"shield_providers,omitempty"`
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
	path := fmt.Sprintf(originsBasePath, origin.Service)
	return Create[Origin](client, path, origin)
}

func (client *IORiverClient) UpdateOrigin(origin Origin) (*Origin, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originsBasePath, origin.Service), origin.Id)
	return Update[Origin](client, path, origin)
}

func (client *IORiverClient) DeleteOrigin(serviceId string, originId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(originsBasePath, serviceId), originId)
	return Delete(client, path)
}
