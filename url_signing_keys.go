package ioriver

import "fmt"

type UrlSigningKey struct {
	Id      string `json:"id,omitempty"`
	Service string `json:"service"`
	Name    string `json:"name"`

	PublicKey     string `json:"public_key"`
	EncryptionKey string `json:"encryption_key"`

	ProviderKeys map[string]string `json:"provider_keys,omitempty"`
}

const urlSginingKeyBasePath = `services/%s/url-signing-keys/`

func (client *IORiverClient) GetUrlSigningKey(serviceId string, id string) (*UrlSigningKey, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(urlSginingKeyBasePath, serviceId), id)
	return Get[UrlSigningKey](client, path)
}

func (client *IORiverClient) ListUrlSigningKeys(serviceId string) ([]UrlSigningKey, error) {
	path := fmt.Sprintf(urlSginingKeyBasePath, serviceId)
	return List[UrlSigningKey](client, path)
}

func (client *IORiverClient) CreateUrlSigningKey(urlSginingKey UrlSigningKey) (*UrlSigningKey, error) {
	path := fmt.Sprintf(urlSginingKeyBasePath, urlSginingKey.Service)
	return Create[UrlSigningKey](client, path, urlSginingKey)
}

func (client *IORiverClient) UpdateUrlSigningKey(urlSginingKey UrlSigningKey) (*UrlSigningKey, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(urlSginingKeyBasePath, urlSginingKey.Service), urlSginingKey.Id)
	return Update[UrlSigningKey](client, path, urlSginingKey)
}

func (client *IORiverClient) DeleteUrlSigningKey(serviceId string, urlSginingKeyId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(urlSginingKeyBasePath, serviceId), urlSginingKeyId)
	return Delete(client, path)
}
