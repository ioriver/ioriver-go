package ioriver

import (
	"fmt"
)

type ActionType string

const (
	ALLOWED_METHODS             ActionType = "ALLOWED_METHODS"
	AUTO_MINIFY                 ActionType = "AUTO_MINIFY"
	BROWSER_CACHE_TTL           ActionType = "BROWSER_CACHE_TTL"
	BYPASS_CACHE_ON_COOKIE      ActionType = "BYPASS_CACHE_ON_COOKIE"
	CACHE_BEHAVIOR              ActionType = "CACHE_BEHAVIOR"
	CACHE_KEY                   ActionType = "CACHE_KEY"
	CACHED_METHODS              ActionType = "CACHED_METHODS"
	CACHE_TTL                   ActionType = "CACHE_TTL"
	COMPRESSION                 ActionType = "COMPRESSION"
	DELETE_RESPONSE_HEADER      ActionType = "DELETE_RESPONSE_HEADER"
	DELETE_REQUEST_HEADER       ActionType = "DELETE_REQUEST_HEADER"
	DENY_ACCESS                 ActionType = "DENY_ACCESS"
	DENY_ACCESS_BY_IP           ActionType = "DENY_ACCESS_BY_IP"
	DENY_ACCESS_BY_TIME         ActionType = "DENY_ACCESS_BY_TIME"
	FOLLOW_REDIRECTS            ActionType = "FOLLOW_REDIRECTS"
	FORWARD_CLIENT_HEADER       ActionType = "FORWARD_CLIENT_HEADER"
	GENERATE_PREFLIGHT_RESPONSE ActionType = "GENERATE_PREFLIGHT_RESPONSE"
	GENERATE_RESPONSE           ActionType = "GENERATE_RESPONSE"
	HOST_HEADER_OVERRIDE        ActionType = "HOST_HEADER_OVERRIDE"
	LARGE_FILES_OPTIMIZATION    ActionType = "LARGE_FILES_OPTIMIZATION"
	ORIGIN_CACHE_CONTROL        ActionType = "ORIGIN_CACHE_CONTROL"
	ORIGIN_ERRORS_PASS_THRU     ActionType = "ORIGIN_ERRORS_PASS_THRU"
	OVERRIDE_ORIGIN             ActionType = "OVERRIDE_ORIGIN"
	REDIRECT                    ActionType = "REDIRECT"
	SET_CORS_HEADER             ActionType = "SET_CORS_HEADER"
	SET_REQUEST_HEADER          ActionType = "SET_REQUEST_HEADER"
	SET_RESPONSE_HEADER         ActionType = "SET_RESPONSE_HEADER"
	STALE_TTL                   ActionType = "STALE_TTL"
	STATUS_CODE_BROWSER_CACHE   ActionType = "STATUS_CODE_BROWSER_CACHE"
	STATUS_CODE_CACHE           ActionType = "STATUS_CODE_CACHE"
	STREAM_LOGS                 ActionType = "STREAM_LOGS"
	TRUE_CLIENT_IP              ActionType = "TRUE_CLIENT_IP"
	URL_REWRITE                 ActionType = "URL_REWRITE"
	URL_SIGNING                 ActionType = "URL_SIGNING"
	VIEWER_PROTOCOL             ActionType = "VIEWER_PROTOCOL"
)

type DateTimeWindowModel struct {
	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`
}

type PeriodicWindowModel struct {
	StartDate         int    `json:"start_date"`
	Duration          int    `json:"duration"`
	DurationUnits     string `json:"duration_units"`
	RepeatPeriod      int    `json:"repeat_period"`
	RepeatPeriodUnits string `json:"repeat_period_units"`
}

type DenyAccessByTimeModel struct {
	DateTimeWindow *DateTimeWindowModel `json:"date_time_window,omitempty"`
	TimePeriodic   *PeriodicWindowModel `json:"time_periodic,omitempty"`
}

type DenyAccessByIPModel struct {
	IPList []string `json:"ip_list"`
}

type BehaviorAction struct {
	Id                        string                 `json:"id,omitempty"`
	Type                      ActionType             `json:"type"`
	MaxTTL                    int                    `json:"max_ttl,omitempty"`
	ResponseHeaderName        string                 `json:"response_header_name,omitempty"`
	ResponseHeaderValue       string                 `json:"response_header_value,omitempty"`
	RequestHeaderName         string                 `json:"request_header_name,omitempty"`
	RequestHeaderValue        string                 `json:"request_header_value,omitempty"`
	CacheBehaviorValue        string                 `json:"cache_behavior_value,omitempty"`
	RedirectURL               string                 `json:"redirect_url,omitempty"`
	OriginCacheControlEnabled bool                   `json:"origin_cache_control_enabled,omitempty"`
	Pattern                   string                 `json:"pattern,omitempty"`
	Cookie                    string                 `json:"cookie,omitempty"`
	AutoMinify                string                 `json:"auto_minify,omitempty"`
	HostHeader                string                 `json:"host_header,omitempty"`
	UseOriginHost             *bool                  `json:"use_domain_origin,omitempty"`
	Origin                    string                 `json:"origin,omitempty"`
	Enabled                   *bool                  `json:"enabled,omitempty"`
	CacheKey                  string                 `json:"cache_key,omitempty"`
	ClientHeaderName          string                 `json:"client_header_name,omitempty"`
	ActionDisabled            bool                   `json:"action_disabled,omitempty"`
	StatusCode                int                    `json:"status_code,omitempty"`
	UnifiedLogDestination     string                 `json:"unified_log_destination,omitempty"`
	UnifiedLogSamplingRate    int                    `json:"unified_log_sampling_rate,omitempty"`
	AllowedMethods            string                 `json:"allowed_methods,omitempty"`
	ResponsePagePath          string                 `json:"response_page_path,omitempty"`
	CachedMethods             string                 `json:"cached_methods,omitempty"`
	DenyAccessByTime          *DenyAccessByTimeModel `json:"deny_access_by_time,omitempty"`
	DenyAccessByIP            *DenyAccessByIPModel   `json:"deny_access_by_ip,omitempty"`
	ViewerProtocol            string                 `json:"viewer_protocol,omitempty"`
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
