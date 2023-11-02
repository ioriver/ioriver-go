package ioriver

import (
	"fmt"
)

type RuleActionType string

const (
	ACTION_BLOCK  RuleActionType = "deny"
	ACTION_LOG    RuleActionType = "log"
	ACTION_BYPASS RuleActionType = "bypass"
)

type RuleType string

const (
	WAF_RULE        RuleType = "WAF"
	RATE_LIMIT_RULE RuleType = "RateLimit"
)

type WAFCustomRuleCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type WAFCustomRuleConditionGroup struct {
	Conditions []WAFCustomRuleCondition `json:"conditions"`
}

type WAFCustomRule struct {
	Id                string                        `json:"id,omitempty"`
	Action            RuleActionType                `json:"action"`
	ConditionGroups   []WAFCustomRuleConditionGroup `json:"condition_groups"`
	Name              string                        `json:"name"`
	Enabled           bool                          `json:"enabled"`
	Type              RuleType                      `json:"type"`
	NumRequests       int                           `json:"num_of_requests,omitempty"`
	TimeWindowSeconds int                           `json:"time_window_seconds,omitempty"`
	DurationSeconds   int                           `json:"duration_seconds,omitempty"`
	Priority          int                           `json:"priority"`
}

const wafCustomRuleBasePath = `services/%s/waf-custom-rules/`

func (client *IORiverClient) GetWAFCustomRule(serviceId string, id string) (*WAFCustomRule, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(wafCustomRuleBasePath, serviceId), id)
	return Get[WAFCustomRule](client, path)
}

func (client *IORiverClient) ListWAFCustomRules(serviceId string) ([]WAFCustomRule, error) {
	path := fmt.Sprintf(wafCustomRuleBasePath, serviceId)
	return List[WAFCustomRule](client, path)
}

func (client *IORiverClient) CreateWAFCustomRule(serviceId string, wafCustomRule WAFCustomRule) (*WAFCustomRule, error) {
	path := fmt.Sprintf(wafCustomRuleBasePath, serviceId)
	return Create[WAFCustomRule](client, path, wafCustomRule)
}

func (client *IORiverClient) UpdateWAFCustomRule(serviceId string, wafCustomRule WAFCustomRule) (*WAFCustomRule, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(wafCustomRuleBasePath, serviceId), wafCustomRule.Id)
	return Update[WAFCustomRule](client, path, wafCustomRule)
}

func (client *IORiverClient) DeleteWAFCustomRule(serviceId string, wafCustomRuleId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(wafCustomRuleBasePath, serviceId), wafCustomRuleId)
	return Delete(client, path)
}
