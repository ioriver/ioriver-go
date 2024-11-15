package ioriver

import (
	"fmt"
)

type ActionType string

const (
	SET_RESPONSE_HEADER         ActionType = "SET_RESPONSE_HEADER"
	CACHE_TTL                   ActionType = "CACHE_TTL"
	REDIRECT_HTTP_TO_HTTPS      ActionType = "REDIRECT_HTTP_TO_HTTPS"
	CACHE_BEHAVIOR              ActionType = "CACHE_BEHAVIOR"
	BROWSER_CACHE_TTL           ActionType = "BROWSER_CACHE_TTL"
	REDIRECT                    ActionType = "REDIRECT"
	ORIGIN_CACHE_CONTROL        ActionType = "ORIGIN_CACHE_CONTROL"
	BYPASS_CACHE_ON_COOKIE      ActionType = "BYPASS_CACHE_ON_COOKIE"
	CACHE_KEY                   ActionType = "CACHE_KEY"
	AUTO_MINIFY                 ActionType = "AUTO_MINIFY"
	HOST_HEADER_OVERRIDE        ActionType = "HOST_HEADER_OVERRIDE"
	SET_CORS_HEADER             ActionType = "SET_CORS_HEADER"
	OVERRIDE_ORIGIN             ActionType = "OVERRIDE_ORIGIN"
	ORIGIN_ERRORS_PASS_THRU     ActionType = "ORIGIN_ERRORS_PASS_THRU"
	FORWARD_CLIENT_HEADER       ActionType = "FORWARD_CLIENT_HEADER"
	FOLLOW_REDIRECTS            ActionType = "FOLLOW_REDIRECTS"
	STATUS_CODE_CACHE           ActionType = "STATUS_CODE_CACHE"
	STATUS_CODE_BROWSER_CACHE   ActionType = "STATUS_CODE_BROWSER_CACHE"
	GENERATE_PREFLIGHT_RESPONSE ActionType = "GENERATE_PREFLIGHT_RESPONSE"
	STALE_TTL                   ActionType = "STALE_TTL"
	STREAM_LOGS                 ActionType = "STREAM_LOGS"
	ALLOWED_METHODS             ActionType = "ALLOWED_METHODS"
	COMPRESSION                 ActionType = "COMPRESSION"
	GENERATE_RESPONSE           ActionType = "GENERATE_RESPONSE"
	CACHED_METHODS              ActionType = "CACHED_METHODS"
)

type BehaviorAction struct {
	Id                        string     `json:"id,omitempty"`
	Type                      ActionType `json:"type"`
	MaxTTL                    int        `json:"max_ttl,omitempty"`
	ResponseHeaderName        string     `json:"response_header_name,omitempty"`
	ResponseHeaderValue       string     `json:"response_header_value,omitempty"`
	CacheBehaviorValue        string     `json:"cache_behavior_value,omitempty"`
	RedirectURL               string     `json:"redirect_url,omitempty"`
	OriginCacheControlEnabled bool       `json:"origin_cache_control_enabled,omitempty"`
	Pattern                   string     `json:"pattern,omitempty"`
	Cookie                    string     `json:"cookie,omitempty"`
	AutoMinify                string     `json:"auto_minify,omitempty"`
	HostHeader                string     `json:"host_header,omitempty"`
	Origin                    string     `json:"origin,omitempty"`
	Enabled                   *bool      `json:"enabled,omitempty"`
	CacheKey                  string     `json:"cache_key,omitempty"`
	ClientHeaderName          string     `json:"client_header_name,omitempty"`
	ActionDisabled            bool       `json:"action_disabled,omitempty"`
	StatusCode                int        `json:"status_code,omitempty"`
	UnifiedLogDestination     string     `json:"unified_log_destination,omitempty"`
	UnifiedLogSamplingRate    int        `json:"unified_log_sampling_rate,omitempty"`
	AllowedMethods            string     `json:"allowed_methods,omitempty"`
	ResponsePagePath          string     `json:"response_page_path,omitempty"`
	CachedMethods             string     `json:"cached_methods,omitempty"`
}

type Behavior struct {
	Id          string           `json:"id,omitempty"`
	Service     string           `json:"service"`
	Name        string           `json:"name"`
	PathPattern string           `json:"path_pattern"`
	Actions     []BehaviorAction `json:"behavior_actions"`
	IsDefault   bool             `json:"is_default,omitempty"`
}

const behaviorsBasePath = `services/%s/behaviors/`

func (client *IORiverClient) GetBehavior(serviceId string, id string) (*Behavior, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(behaviorsBasePath, serviceId), id)
	return Get[Behavior](client, path)
}

func (client *IORiverClient) ListBehaviors(serviceId string) ([]Behavior, error) {
	path := fmt.Sprintf(behaviorsBasePath, serviceId)
	return List[Behavior](client, path)
}

func (client *IORiverClient) CreateBehavior(behavior Behavior) (*Behavior, error) {
	path := fmt.Sprintf(behaviorsBasePath, behavior.Service)
	return Create[Behavior](client, path, behavior)
}

func (client *IORiverClient) UpdateBehavior(behavior Behavior) (*Behavior, error) {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(behaviorsBasePath, behavior.Service), behavior.Id)
	return Update[Behavior](client, path, behavior)
}

func (client *IORiverClient) DeleteBehavior(serviceId string, behaviorId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(behaviorsBasePath, serviceId), behaviorId)
	return Delete(client, path)
}

func (client *IORiverClient) ResetDefaultBehavior(serviceId string) error {
	path := fmt.Sprintf("%s%s/", fmt.Sprintf(behaviorsBasePath, serviceId), "reset_default_behavior")
	_, err := Update(client, path, Behavior{})
	return err
}
