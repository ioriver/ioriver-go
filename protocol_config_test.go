package ioriver

import (
	"fmt"
	"testing"
)

const (
	testProtocolConfigHttp2Value = true
	testProtocolConfigHttp3Value = false
	testProtocolConfigIpv6Value  = false
)

const serverProtocolConfigData = `{
	"id":"%s",
	"service":"%s",
	"http2_enabled":%t,
	"http3_enabled":%t,
	"ipv6_enabled":%t
}`

var expectedProtocolConfig = ProtocolConfig{
	Id:           testObjectId,
	Service:      testServiceId,
	Http2Enabled: true,
	Http3Enabled: false,
}

func TestListProtocolConfigs(t *testing.T) {
	path := fmt.Sprintf("/services/%s/protocol-config/", testServiceId)
	serverData := fmt.Sprintf(`[%s]`, fmt.Sprintf(serverProtocolConfigData, testObjectId, testServiceId,
		testProtocolConfigHttp2Value, testProtocolConfigHttp3Value, testProtocolConfigIpv6Value))
	expected := []ProtocolConfig{expectedProtocolConfig}
	RunServiceList[ProtocolConfig](t, (*IORiverClient).ListProtocolConfigs, path, testServiceId, serverData, expected)
}

func TestGetProtocolConfig(t *testing.T) {
	path := fmt.Sprintf("/services/%s/protocol-config/%s/", testServiceId, testObjectId)
	serverData := fmt.Sprintf(serverProtocolConfigData, testObjectId, testServiceId, testProtocolConfigHttp2Value,
		testProtocolConfigHttp3Value, testProtocolConfigIpv6Value)
	RunServiceGet[ProtocolConfig](t, (*IORiverClient).GetProtocolConfig, path, testServiceId, testObjectId, serverData,
		&expectedProtocolConfig)
}

func TestCreateProtocolConfig(t *testing.T) {
	newProtocolConfig := ProtocolConfig{
		Service:      testServiceId,
		Http2Enabled: testProtocolConfigHttp2Value,
		Http3Enabled: testProtocolConfigHttp3Value,
		Ipv6Enabled:  testProtocolConfigIpv6Value,
	}

	path := fmt.Sprintf("/services/%s/protocol-config/", testServiceId)
	serverData := fmt.Sprintf(serverProtocolConfigData, testObjectId, testServiceId, testProtocolConfigHttp2Value,
		testProtocolConfigHttp3Value, testProtocolConfigIpv6Value)
	RunCreate[ProtocolConfig](t, (*IORiverClient).CreateProtocolConfig, path, newProtocolConfig, serverData,
		&expectedProtocolConfig)
}
