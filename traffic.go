package ioriver

import "fmt"

type Granularity int

const (
	Minute Granularity = iota
	Hour
	Day
)

var granularityName = map[Granularity]string{
	Minute: "MINUTE",
	Hour:   "HOUR",
	Day:    "DAY",
}

func (granularity Granularity) String() string {
	return granularityName[granularity]
}

type AdvancedMetric int

const (
	StatusCode AdvancedMetric = iota
	HttpVersion
	HttpMethod
)

var advancedMetricName = map[AdvancedMetric]string{
	StatusCode:  "status_code",
	HttpVersion: "http_version",
	HttpMethod:  "method",
}

func (advancedMetric AdvancedMetric) String() string {
	return advancedMetricName[advancedMetric]
}

type Point struct {
	Timestamp int64    `json:"timestamp"`
	Metrics   []Metric `json:"metrics"`
}

type Metric struct {
	AdvancedMetricName  *string `json:"advancedMetricName"`
	AdvancedMetricValue *string `json:"advancedMetricValue"`
	Geo                 *string `json:"geo"`
	Metrics             Metrics `json:"metrics"`
	ProviderName        string  `json:"providerName"`
}

type Metrics struct {
	Bytes                     int     `json:"bytes"`
	CachedBytesPercentage     float64 `json:"cachedBytesPercentage"`
	CachedHitsPercentage      float64 `json:"cachedHitsPercentage"`
	EdgeCachedBytesPercentage float64 `json:"edgeCachedBytesPercentage"`
	EdgeCachedHitsPercentage  float64 `json:"edgeCachedHitsPercentage"`
	ErrorsPercentage          float64 `json:"errorsPercentage"`
	Hits                      int     `json:"hits"`
	MidgressBytes             int     `json:"midgressBytes"`
	MidgressHits              int     `json:"midgressHits"`
	NumMinutes                int     `json:"numMinutes"`
	OriginBytes               int     `json:"originBytes"`
	OriginHits                int     `json:"originHits"`
}

type ServiceStats struct {
	ServiceId string  `json:"serviceID"`
	Points    []Point `json:"points"`
}

type Traffic struct {
	ServiceStats []ServiceStats `json:"serviceStats"`
	Granularity  string         `json:"granularity"`
	Error        string         `json:"error"`
}

const trafficBasePath = "traffic/overtime/"

func (client *IORiverClient) GetTraffic(serviceId string, startTime int64, endTime int64, granularity Granularity) (*Traffic, error) {
	startTimeParam, endTimeParam, granularityParam := getQueryParams(startTime, endTime, granularity)
	path := fmt.Sprintf("%s%s/", trafficBasePath, serviceId)
	return Get[Traffic](client, path, startTimeParam, endTimeParam, granularityParam)
}

func (client *IORiverClient) GetAdvancedTraffic(
	serviceId string, startTime int64, endTime int64, granularity Granularity, advancedMetric AdvancedMetric) (*Traffic, error) {
	path := fmt.Sprintf("%s%s/", trafficBasePath, serviceId)
	startTimeParam, endTimeParam, granularityParam := getQueryParams(startTime, endTime, granularity)
	advancedMetricParams := fmt.Sprintf("advancedMetricName=%s", advancedMetric)
	return Get[Traffic](client, path, startTimeParam, endTimeParam, granularityParam, advancedMetricParams)
}

func getQueryParams(startTime int64, endTime int64, granularity Granularity) (string, string, string) {
	startTimeParam := fmt.Sprintf("startTime=%d", startTime)
	endTimeParam := fmt.Sprintf("endTime=%d", endTime)
	granularityParam := fmt.Sprintf("granularity=%s", granularity)
	return startTimeParam, endTimeParam, granularityParam
}

func (t *Traffic) GetFilteredMetrics(serviceId string, predicate func(metric *Metric, timestamp int64) bool) []Metric {
	filtered := []Metric{}

	for _, stat := range t.ServiceStats {
		if stat.ServiceId == serviceId {
			for _, point := range stat.Points {
				for _, metrics := range point.Metrics {
					if predicate(&metrics, point.Timestamp) {
						filtered = append(filtered, metrics)
					}
				}
			}
		}
	}
	return filtered
}
