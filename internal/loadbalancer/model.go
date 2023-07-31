package loadbalancer

// MockLoadBalancerData represents the load balancer data
type MockLoadBalancerData struct {
	CPUUtilization    float64           `json:"cpu_utilization"`
	MemoryUsage       float64           `json:"memory_usage"`
	NetworkThroughput NetworkThroughput `json:"network_throughput"`
	RequestLatencyMS  float64           `json:"request_latency_ms"`
	ErrorRate         float64           `json:"error_rate"`
	ActiveConnections int               `json:"active_connections"`
	BackendServers    int               `json:"backend_servers"`
	Algorithm         string            `json:"algorithm"`
	Uptime            string            `json:"uptime"`
}

// NetworkThroughput represents network throughput data corresponding to load balancer
type NetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}
