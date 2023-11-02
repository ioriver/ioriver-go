package ioriver

import (
	"fmt"
)

type WAF struct {
	Id      string `json:"id,omitempty"`
	Service string `json:"service"`
	Enabled bool   `json:"enabled"`
}

const wafBasePath = `services/%s/waf/`

func (client *IORiverClient) GetWAF(serviceId string) (*WAF, error) {
	path := fmt.Sprintf(wafBasePath, serviceId)
	return Get[WAF](client, path)
}

func (client *IORiverClient) UpdateWAF(waf WAF) (*WAF, error) {
	path := fmt.Sprintf(wafBasePath, waf.Service)
	return Update[WAF](client, path, waf)
}
