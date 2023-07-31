package linuxserver

// MockLinuxServerData represents the linux server data
type MockLinuxServerData struct {
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

// NetworkThroughput represents the network throughput corresponding to linux server
type NetworkThroughput struct {
	Sent     int `json:"sent"`
	Received int `json:"received"`
}

// DiskUsage represents the disk usage corresponding to linux server
type DiskUsage struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

// LoadAverage represents the load average data corresponding to linux server
type LoadAverage struct {
	Min1  float64 `json:"1min"`
	Min5  float64 `json:"5min"`
	Min15 float64 `json:"15min"`
}
