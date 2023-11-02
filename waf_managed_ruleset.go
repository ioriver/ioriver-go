package ioriver

import (
	"fmt"
)

type WAFManagedRuleGroup struct {
	Id         string `json:"id,omitempty"`
	RulePrefix string `json:"rule_prefix,omitempty"`
	Enabled    bool   `json:"enabled"`
}

type WAFOverrideManagedRule struct {
	Id      string `json:"id,omitempty"`
	RuleId  string `json:"rule_id"`
	Enabled bool   `json:"enabled"`
}

type WAFManagedRuleset struct {
	Id               string                   `json:"id,omitempty"`
	Name             string                   `json:"name"`
	DisplayName      string                   `json:"display_name"`
	Enabled          bool                     `json:"enabled"`
	Block            bool                     `json:"block,omitempty"`
	ParanoiaLevel    int                      `json:"paranoia_level,omitempty"`
	AnomalyThreshold int                      `json:"anomaly_threshold,omitempty"`
	Groups           []WAFManagedRuleGroup    `json:"groups"`
	Overrides        []WAFOverrideManagedRule `json:"overrides"`
}

const wafManagedRsBasePath = `services/%s/waf-managed-rulesets/`

func (client *IORiverClient) GetWAFManagedRuleset(serviceId string, id string) (*WAFManagedRuleset, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(wafManagedRsBasePath, serviceId), id)
	return Get[WAFManagedRuleset](client, path)
}

func (client *IORiverClient) ListWAFManagedRuleset(serviceId string) ([]WAFManagedRuleset, error) {
	path := fmt.Sprintf(wafManagedRsBasePath, serviceId)
	return List[WAFManagedRuleset](client, path)
}

func (client *IORiverClient) UpdateWAFManagedRuleset(serviceId string, wafManagedRuleset WAFManagedRuleset) (*WAFManagedRuleset, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(wafManagedRsBasePath, serviceId), wafManagedRuleset.Id)
	return Update[WAFManagedRuleset](client, path, wafManagedRuleset)
}
