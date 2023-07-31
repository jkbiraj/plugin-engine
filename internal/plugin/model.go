package plugin

// FirewallDataResponse represents the firewall data response
type FirewallDataResponse struct {
	CPUUtilization    float64                   `json:"cpu_utilization"`
	MemoryUsage       float64                   `json:"memory_usage"`
	NetworkThroughput FireWallNetworkThroughput `json:"network_throughput"`
	RequestLatencyMS  float64                   `json:"request_latency_ms"`
	ErrorRate         float64                   `json:"error_rate"`
	PacketsFiltered   int                       `json:"packets_filtered"`
	ActiveConnections int                       `json:"active_connections"`
	Rules             int                       `json:"rules"`
	Uptime            string                    `json:"uptime"`
	Interfaces        map[string]InterfaceStats `json:"interfaces"`
}

// FireWallNetworkThroughput represents the firewall network throughput
type FireWallNetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// InterfaceStats represents the interface statistics data
type InterfaceStats struct {
	IPAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
	PacketsIn  int    `json:"packets_in"`
	PacketsOut int    `json:"packets_out"`
}

// LinuxServerDataResponse represents the linux server data response
type LinuxServerDataResponse struct {
	CPUUtilization    float64           `json:"cpu_utilization"`
	MemoryUsage       float64           `json:"memory_usage"`
	NetworkThroughput NetworkThroughput `json:"network_throughput"`
	RequestLatencyMS  float64           `json:"request_latency_ms"`
	ErrorRate         float64           `json:"error_rate"`
	DiskUsage         DiskUsage         `json:"disk_usage"`
	ActiveConnections int               `json:"active_connections"`
	Processes         int               `json:"processes"`
	Uptime            string            `json:"uptime"`
	LoadAverage       LoadAverage       `json:"load_average"`
}

// NetworkThroughput represents the network throughput data
type NetworkThroughput struct {
	Sent     int `json:"sent"`
	Received int `json:"received"`
}

// DiskUsage represents the disk usage data
type DiskUsage struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

// LoadAverage represents the average load data
type LoadAverage struct {
	Min1  float64 `json:"1min"`
	Min5  float64 `json:"5min"`
	Min15 float64 `json:"15min"`
}

// LoadBalancerDataResponse represents the load balancer data response
type LoadBalancerDataResponse struct {
	CPUUtilization    float64                       `json:"cpu_utilization"`
	MemoryUsage       float64                       `json:"memory_usage"`
	NetworkThroughput LoadBalancerNetworkThroughput `json:"network_throughput"`
	RequestLatencyMS  float64                       `json:"request_latency_ms"`
	ErrorRate         float64                       `json:"error_rate"`
	ActiveConnections int                           `json:"active_connections"`
	BackendServers    int                           `json:"backend_servers"`
	Algorithm         string                        `json:"algorithm"`
	Uptime            string                        `json:"uptime"`
}

// LoadBalancerNetworkThroughput represents the load balancer network throughput
type LoadBalancerNetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// RouterDataResponse represents the router data response
type RouterDataResponse struct {
	RouterID          string                  `json:"router_id"`
	CPUUtilization    CPUUtilization          `json:"cpu_utilization"`
	MemoryUsage       MemoryUsage             `json:"memory_usage"`
	NetworkThroughput RouterNetworkThroughput `json:"network_throughput"`
	RequestLatency    RequestLatency          `json:"request_latency"`
	ErrorRates        ErrorRates              `json:"error_rates"`
	OperationalParams OperationalParams       `json:"operational_parameters"`
}

// CPUUtilization represents the cpu utilization data
type CPUUtilization struct {
	Average float64 `json:"average"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

// MemoryUsage represents the memory usage data
type MemoryUsage struct {
	Used       int     `json:"used"`
	Total      int     `json:"total"`
	Percentage float64 `json:"percentage"`
}

// RouterNetworkThroughput represents the router network throughput data
type RouterNetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// RequestLatency represents the request latency data
type RequestLatency struct {
	AverageMs float64 `json:"average_ms"`
	MaxMs     float64 `json:"max_ms"`
}

// ErrorRates represents the error rates data
type ErrorRates struct {
	HTTP500 float64 `json:"http_500"`
	HTTP404 float64 `json:"http_404"`
}

// OperationalParams represents the other operational params
type OperationalParams struct {
	IsActive    bool    `json:"is_active"`
	Uptime      string  `json:"uptime"`
	Temperature float64 `json:"temperature"`
}

// SwitchDataResponse represents the switch data response
type SwitchDataResponse struct {
	SwitchID          string                  `json:"switch_id"`
	CPUUtilization    CPUUtilization          `json:"cpu_utilization"`
	MemoryUsage       MemoryUsage             `json:"memory_usage"`
	NetworkThroughput SwitchNetworkThroughput `json:"network_throughput"`
	RequestLatency    RequestLatency          `json:"request_latency"`
	ErrorRates        SwitchErrorRates        `json:"error_rates"`
	OperationalParams OperationalParams       `json:"operational_parameters"`
}

// SwitchNetworkThroughput represents the switch network throughput
type SwitchNetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// SwitchErrorRates represents the switch error rates
type SwitchErrorRates struct {
	TCPConnection float64 `json:"tcp_connection"`
	UDPPacketLoss float64 `json:"udp_packet_loss"`
}

// WindowsServerDataResponse represents the Windows server data response
type WindowsServerDataResponse struct {
	ServerID          string                         `json:"server_id"`
	CPUUtilization    CPUUtilization                 `json:"cpu_utilization"`
	MemoryUsage       MemoryUsage                    `json:"memory_usage"`
	NetworkThroughput WindowsServerNetworkThroughput `json:"network_throughput"`
	RequestLatency    RequestLatency                 `json:"request_latency"`
	ErrorRates        ErrorRates                     `json:"error_rates"`
	OperationalParams WindowsServerOperationalParams `json:"operational_parameters"`
}

// WindowsServerNetworkThroughput represents the Windows server network throughput
type WindowsServerNetworkThroughput struct {
	Incoming float64 `json:"incoming"`
	Outgoing float64 `json:"outgoing"`
}

// WindowsServerOperationalParams represents other Windows server operational params
type WindowsServerOperationalParams struct {
	IsOnline    bool    `json:"is_online"`
	Uptime      string  `json:"uptime"`
	Temperature float64 `json:"temperature"`
}
