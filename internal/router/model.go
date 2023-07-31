package router

// MockRouterData represents the router data
type MockRouterData struct {
	RouterID          string            `json:"router_id"`
	CPUUtilization    CPUUtilization    `json:"cpu_utilization"`
	MemoryUsage       MemoryUsage       `json:"memory_usage"`
	NetworkThroughput NetworkThroughput `json:"network_throughput"`
	RequestLatency    RequestLatency    `json:"request_latency"`
	ErrorRates        ErrorRates        `json:"error_rates"`
	OperationalParams OperationalParams `json:"operational_parameters"`
}

// CPUUtilization represents the cpu utilization data corresponding to router
type CPUUtilization struct {
	Average float64 `json:"average"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

// MemoryUsage represents the memory usage data corresponding to router
type MemoryUsage struct {
	Used       int     `json:"used"`
	Total      int     `json:"total"`
	Percentage float64 `json:"percentage"`
}

// NetworkThroughput represents the network throughput data corresponding to router
type NetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// RequestLatency represents the request latency data corresponding to router
type RequestLatency struct {
	AverageMs float64 `json:"average_ms"`
	MaxMs     float64 `json:"max_ms"`
}

// ErrorRates represents the error rate data corresponding to router
type ErrorRates struct {
	HTTP500 float64 `json:"http_500"`
	HTTP404 float64 `json:"http_404"`
}

// OperationalParams represents the other operational data corresponding to router
type OperationalParams struct {
	IsActive    bool    `json:"is_active"`
	Uptime      string  `json:"uptime"`
	Temperature float64 `json:"temperature"`
}
