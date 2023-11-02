package ioriver

import (
	"fmt"
	"testing"
)

const (
	testAccountProviderId   = "45901552-c131-4e94-be27-60267f54d311"
	testProviderId          = 3
	testProviderDetailsId   = testProviderId
	testProviderDetailsName = "Cloudfront"
)

const serverAccountProviderData = `{
	"id":"%s",
	"provider":%d,
	"provider_details": {
		"id":%d,
		"name":"%s"
	}
}
`

var expectedAccoountProvider = AccountProvider{
	Id:       testAccountProviderId,
	Provider: testProviderId,
	Details: ProviderDetails{
		Id:   3,
		Name: "Cloudfront",
	},
}

func TestListAccountProviders(t *testing.T) {
	path := "/account-providers/"
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverAccountProviderData, testAccountProviderId, testProviderId,
		testProviderDetailsId, testProviderDetailsName))
	expected := []AccountProvider{expectedAccoountProvider}
	RunList[AccountProvider](t, (*IORiverClient).ListAccountProviders, path, serverData, expected)
}

func TestGetAccountProvider(t *testing.T) {
	path := fmt.Sprintf("/account-providers/%s/", testAccountProviderId)
	serverData := fmt.Sprintf(serverAccountProviderData, testAccountProviderId, testProviderId,
		testProviderDetailsId, testProviderDetailsName)
	RunGet[AccountProvider](t, (*IORiverClient).GetAccountProvider, path, testAccountProviderId, serverData, &expectedAccoountProvider)
}

func TestCreateAccountProvider(t *testing.T) {
	newProvider := AccountProvider{
		Id:       testAccountProviderId,
		Provider: testProviderId,
	}

	path := "/account-providers/"
	serverData := fmt.Sprintf(serverAccountProviderData, testAccountProviderId, testProviderId,
		testProviderDetailsId, testProviderDetailsName)
	RunCreate[AccountProvider](t, (*IORiverClient).CreateAccountProvider, path, newProvider, serverData, &expectedAccoountProvider)
}
