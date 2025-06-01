package ioriver

import (
	"fmt"
	"testing"
)

const (
	testUrlSigningKeyName = "test-su"
)

const serverUrlSigningKeyData = `{
	"id":"%s",
	"service":"%s",
	"name":"%s",
	"provider_keys": {
			"Fastly": "746PTXZ5DF",
			"Cloudfront": "JU3P6REILPJ8N4"
	}
}`

var expectedUrlSigningKey = UrlSigningKey{
	Id:      testObjectId,
	Service: testServiceId,
	Name:    testUrlSigningKeyName,
	ProviderKeys: map[string]string{
		"Fastly":     "746PTXZ5DF",
		"Cloudfront": "JU3P6REILPJ8N4",
	},
}

func TestListUrlSigningKeys(t *testing.T) {
	path := fmt.Sprintf("/services/%s/url-signing-keys/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverUrlSigningKeyData, testObjectId, testServiceId,
		testUrlSigningKeyName))
	expected := []UrlSigningKey{expectedUrlSigningKey}
	RunServiceList[UrlSigningKey](t, (*IORiverClient).ListUrlSigningKeys, path, testServiceId, serverData, expected)
}

func TestGetUrlSigningKey(t *testing.T) {
	path := fmt.Sprintf("/services/%s/url-signing-keys/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverUrlSigningKeyData, testObjectId, testServiceId,
		testUrlSigningKeyName)
	RunServiceGet[UrlSigningKey](t, (*IORiverClient).GetUrlSigningKey, path, testServiceId, testObjectId, serverData,
		&expectedUrlSigningKey)
}

func TestCreateUrlSigningKey(t *testing.T) {
	newUrlSigningKey := UrlSigningKey{
		Service:       testServiceId,
		Name:          testUrlSigningKeyName,
		PublicKey:     "abcd",
		EncryptionKey: "1234",
	}

	path := fmt.Sprintf("/services/%s/url-signing-keys/", testServiceId)
	serverData := fmt.Sprintf(serverUrlSigningKeyData, testObjectId, testServiceId,
		testUrlSigningKeyName)
	RunCreate[UrlSigningKey](t, (*IORiverClient).CreateUrlSigningKey, path, newUrlSigningKey, serverData,
		&expectedUrlSigningKey)
}
