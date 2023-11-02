package ioriver

import (
	"fmt"
)

const (
	Cloudflare int = 2
	Cloudfront int = 3
	AzureCDN   int = 4
	Akamai     int = 5
	Fastly     int = 13
)

type ProviderDetails struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AccountProvider struct {
	Id          string          `json:"id,omitempty"`
	Provider    int             `json:"provider"`
	Credentials interface{}     `json:"credentials,omitempty"`
	Details     ProviderDetails `json:"provider_details,omitempty"`
}

type FastlyCredentials string
type CloudflareCredentials string

type CloudfrontCredentials struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type AzureCdnCredentials struct {
	SubscriptionId    string `json:"subscriptionId"`
	ClientId          string `json:"clientId"`
	TenantId          string `json:"tenantId"`
	ClientSecret      string `json:"clientSecret"`
	ResourceGroupName string `json:"resourceGroupName"`
}

const acBasePath = "account-providers/"

func (client *IORiverClient) GetAccountProvider(id string) (*AccountProvider, error) {
	path := fmt.Sprintf("%s%s/", acBasePath, id)
	return Get[AccountProvider](client, path)
}

func (client *IORiverClient) ListAccountProviders() ([]AccountProvider, error) {
	return List[AccountProvider](client, acBasePath)
}

func (client *IORiverClient) CreateAccountProvider(provider AccountProvider) (*AccountProvider, error) {
	return Create[AccountProvider](client, acBasePath, provider)
}

func (client *IORiverClient) UpdateAccountProvider(provider AccountProvider) (*AccountProvider, error) {
	path := acBasePath + provider.Id + "/"
	return Update[AccountProvider](client, path, provider)
}

func (client *IORiverClient) DeleteAccountProvider(providerId string) error {
	path := fmt.Sprintf("%s%s/", acBasePath, providerId)
	return Delete(client, path)
}
