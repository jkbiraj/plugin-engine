package firewall

// MockFirewallData represents the Firewall data
type MockFirewallData struct {
	CPUUtilization    float64                   `json:"cpu_utilization"`
	MemoryUsage       float64                   `json:"memory_usage"`
	NetworkThroughput NetworkThroughput         `json:"network_throughput"`
	RequestLatencyMS  float64                   `json:"request_latency_ms"`
	ErrorRate         float64                   `json:"error_rate"`
	PacketsFiltered   int                       `json:"packets_filtered"`
	ActiveConnections int                       `json:"active_connections"`
	Rules             int                       `json:"rules"`
	Uptime            string                    `json:"uptime"`
	Interfaces        map[string]InterfaceStats `json:"interfaces"`
}

// NetworkThroughput represents the network throughput corresponding to firewall
type NetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// InterfaceStats represents the interface statistics corresponding to firewall
type InterfaceStats struct {
	IPAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
	PacketsIn  int    `json:"packets_in"`
	PacketsOut int    `json:"packets_out"`
}
