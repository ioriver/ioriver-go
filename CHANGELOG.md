# Changelog

## v1.2.0

### New Features

- **`RetrieveTrafficOvertime`**: Added a new traffic API method that retrieves overtime traffic statistics for multiple services in a single request. Unlike `GetTraffic` (single service, GET), this method sends a POST request with a `TrafficOvertimeRequest` payload containing `serviceIds` and `advancedMetricNames`, while still accepting time range and granularity as query parameters.

- **`TrafficOvertimeRequest`**: New request struct used as the body for the `RetrieveTrafficOvertime` call.

### Improvements

- **`Create` helper**: Now accepts optional variadic query parameters (`queryParams ...string`), which are joined with `&` and appended to the request URL. This enables POST-based endpoints that also require URL query parameters (e.g., `startTime`, `endTime`, `granularity`).
