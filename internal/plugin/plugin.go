package plugin

import (
	"pluggin-engine/internal/firewall"
	"pluggin-engine/internal/linuxserver"
	"pluggin-engine/internal/loadbalancer"
	"pluggin-engine/internal/router"
	_switch "pluggin-engine/internal/switch"
	"pluggin-engine/internal/windowsserver"
)

// Service is the service contract for plugin
//
//go:generate mockgen -destination mocks/plugin.go -package mocks -source plugin.go
type Service interface {
	FetchFirewallData() (FirewallDataResponse, error)
	FetchLinuxServerData() (LinuxServerDataResponse, error)
	FetchLoadBalancerData() (LoadBalancerDataResponse, error)
	FetchRouterData() (RouterDataResponse, error)
	FetchSwitchData() (SwitchDataResponse, error)
	FetchWindowsServerData() (WindowsServerDataResponse, error)
}

// service implements the domain.Service contract
type service struct {
	config               Config
	fireWallService      firewall.Service
	linuxServerService   linuxserver.Service
	loadBalancerService  loadbalancer.Service
	routerService        router.Service
	switchService        _switch.Service
	windowsServerService windowsserver.Service
}

// contract validation
var (
	// _ is used to ensure interface is satisfied
	_ Service = &service{}
)

// NewService returns the service that satisfies the Service interface
func NewService(config Config, fireWallService firewall.Service, linuxServerService linuxserver.Service,
	loadBalancerService loadbalancer.Service, routerService router.Service, switchService _switch.Service,
	windowsServerService windowsserver.Service) Service {
	return &service{
		config:               config,
		fireWallService:      fireWallService,
		linuxServerService:   linuxServerService,
		loadBalancerService:  loadBalancerService,
		routerService:        routerService,
		switchService:        switchService,
		windowsServerService: windowsServerService,
	}
}

// FetchFirewallData returns firewall data
func (s *service) FetchFirewallData() (FirewallDataResponse, error) {
	firewallData, err := s.fireWallService.GetFirewallData()
	if err != nil {
		return FirewallDataResponse{}, err
	}
	interfaces := make(map[string]InterfaceStats, len(firewallData.Interfaces))
	for interfaceKey, stats := range firewallData.Interfaces {
		interfaces[interfaceKey] = InterfaceStats{
			IPAddress:  stats.IPAddress,
			SubnetMask: stats.SubnetMask,
			PacketsIn:  stats.PacketsIn,
			PacketsOut: stats.PacketsOut,
		}
	}
	return FirewallDataResponse{
		CPUUtilization: firewallData.CPUUtilization,
		MemoryUsage:    firewallData.MemoryUsage,
		NetworkThroughput: FireWallNetworkThroughput{
			Incoming: firewallData.NetworkThroughput.Incoming,
			Outgoing: firewallData.NetworkThroughput.Outgoing,
		},
		RequestLatencyMS:  firewallData.RequestLatencyMS,
		ErrorRate:         firewallData.ErrorRate,
		PacketsFiltered:   firewallData.PacketsFiltered,
		ActiveConnections: firewallData.ActiveConnections,
		Rules:             firewallData.Rules,
		Uptime:            firewallData.Uptime,
		Interfaces:        interfaces,
	}, nil
}

// FetchLinuxServerData returns linux server data
func (s *service) FetchLinuxServerData() (LinuxServerDataResponse, error) {
	linuxServerData, err := s.linuxServerService.GetLinuxServerData()
	if err != nil {
		return LinuxServerDataResponse{}, err
	}
	return LinuxServerDataResponse{
		CPUUtilization: linuxServerData.CPUUtilization,
		MemoryUsage:    linuxServerData.MemoryUsage,
		NetworkThroughput: NetworkThroughput{
			Sent:     linuxServerData.NetworkThroughput.Sent,
			Received: linuxServerData.NetworkThroughput.Received,
		},
		RequestLatencyMS: linuxServerData.RequestLatencyMS,
		ErrorRate:        linuxServerData.ErrorRate,
		DiskUsage: DiskUsage{
			Total:     linuxServerData.DiskUsage.Total,
			Used:      linuxServerData.DiskUsage.Used,
			Available: linuxServerData.DiskUsage.Available,
		},
		ActiveConnections: linuxServerData.ActiveConnections,
		Processes:         linuxServerData.Processes,
		Uptime:            linuxServerData.Uptime,
		LoadAverage: LoadAverage{
			Min1:  linuxServerData.LoadAverage.Min1,
			Min5:  linuxServerData.LoadAverage.Min5,
			Min15: linuxServerData.LoadAverage.Min15,
		},
	}, nil
}

// FetchLoadBalancerData returns the load balancer data
func (s *service) FetchLoadBalancerData() (LoadBalancerDataResponse, error) {
	loadBalancerData, err := s.loadBalancerService.GetLoadBalancerData()
	if err != nil {
		return LoadBalancerDataResponse{}, err
	}
	return LoadBalancerDataResponse{
		CPUUtilization: loadBalancerData.CPUUtilization,
		MemoryUsage:    loadBalancerData.MemoryUsage,
		NetworkThroughput: LoadBalancerNetworkThroughput{
			Incoming: loadBalancerData.NetworkThroughput.Incoming,
			Outgoing: loadBalancerData.NetworkThroughput.Outgoing,
		},
		RequestLatencyMS:  loadBalancerData.RequestLatencyMS,
		ErrorRate:         loadBalancerData.ErrorRate,
		ActiveConnections: loadBalancerData.ActiveConnections,
		BackendServers:    loadBalancerData.BackendServers,
		Algorithm:         loadBalancerData.Algorithm,
		Uptime:            loadBalancerData.Uptime,
	}, nil
}

// FetchRouterData returns router data
func (s *service) FetchRouterData() (RouterDataResponse, error) {
	routerData, err := s.routerService.GetRouterData()
	if err != nil {
		return RouterDataResponse{}, err
	}
	return RouterDataResponse{
		RouterID: routerData.RouterID,
		CPUUtilization: CPUUtilization{
			Average: routerData.CPUUtilization.Average,
			Max:     routerData.CPUUtilization.Max,
			Min:     routerData.CPUUtilization.Min,
		},
		MemoryUsage: MemoryUsage{
			Used:       routerData.MemoryUsage.Used,
			Total:      routerData.MemoryUsage.Total,
			Percentage: routerData.MemoryUsage.Percentage,
		},
		NetworkThroughput: RouterNetworkThroughput{
			Incoming: routerData.NetworkThroughput.Incoming,
			Outgoing: routerData.NetworkThroughput.Outgoing,
		},
		RequestLatency: RequestLatency{
			AverageMs: routerData.RequestLatency.AverageMs,
			MaxMs:     routerData.RequestLatency.MaxMs,
		},
		ErrorRates: ErrorRates{
			HTTP500: routerData.ErrorRates.HTTP500,
			HTTP404: routerData.ErrorRates.HTTP404,
		},
		OperationalParams: OperationalParams{
			IsActive:    routerData.OperationalParams.IsActive,
			Uptime:      routerData.OperationalParams.Uptime,
			Temperature: routerData.OperationalParams.Temperature,
		},
	}, nil
}

// FetchSwitchData returns the switch data
func (s *service) FetchSwitchData() (SwitchDataResponse, error) {
	switchData, err := s.switchService.GetSwitchData()
	if err != nil {
		return SwitchDataResponse{}, err
	}
	return SwitchDataResponse{
		SwitchID: switchData.SwitchID,
		CPUUtilization: CPUUtilization{
			Average: switchData.CPUUtilization.Average,
			Max:     switchData.CPUUtilization.Max,
			Min:     switchData.CPUUtilization.Min,
		},
		MemoryUsage: MemoryUsage{
			Used:       switchData.MemoryUsage.Used,
			Total:      switchData.MemoryUsage.Total,
			Percentage: switchData.MemoryUsage.Percentage,
		},
		NetworkThroughput: SwitchNetworkThroughput{
			Incoming: switchData.NetworkThroughput.Incoming,
			Outgoing: switchData.NetworkThroughput.Outgoing,
		},
		RequestLatency: RequestLatency{
			AverageMs: switchData.RequestLatency.AverageMs,
			MaxMs:     switchData.RequestLatency.MaxMs,
		},
		ErrorRates: SwitchErrorRates{
			TCPConnection: switchData.ErrorRates.TCPConnection,
			UDPPacketLoss: switchData.ErrorRates.UDPPacketLoss,
		},
		OperationalParams: OperationalParams{
			IsActive:    switchData.OperationalParams.IsActive,
			Uptime:      switchData.OperationalParams.Uptime,
			Temperature: switchData.OperationalParams.Temperature,
		},
	}, nil
}

// FetchWindowsServerData returns Windows server data
func (s *service) FetchWindowsServerData() (WindowsServerDataResponse, error) {
	windowsServerData, err := s.windowsServerService.GetWindowsServerData()
	if err != nil {
		return WindowsServerDataResponse{}, err
	}

	return WindowsServerDataResponse{
		ServerID: windowsServerData.ServerID,
		CPUUtilization: CPUUtilization{
			Average: windowsServerData.CPUUtilization.Average,
			Max:     windowsServerData.CPUUtilization.Max,
			Min:     windowsServerData.CPUUtilization.Min,
		},
		MemoryUsage: MemoryUsage{
			Used:       windowsServerData.MemoryUsage.Used,
			Total:      windowsServerData.MemoryUsage.Total,
			Percentage: windowsServerData.MemoryUsage.Percentage,
		},
		NetworkThroughput: WindowsServerNetworkThroughput{
			Incoming: windowsServerData.NetworkThroughput.Incoming,
			Outgoing: windowsServerData.NetworkThroughput.Outgoing,
		},
		RequestLatency: RequestLatency{
			AverageMs: windowsServerData.RequestLatency.AverageMs,
			MaxMs:     windowsServerData.RequestLatency.MaxMs,
		},
		ErrorRates: ErrorRates{
			HTTP500: windowsServerData.ErrorRates.HTTP500,
			HTTP404: windowsServerData.ErrorRates.HTTP404,
		},
		OperationalParams: WindowsServerOperationalParams{
			IsOnline:    windowsServerData.OperationalParams.IsOnline,
			Uptime:      windowsServerData.OperationalParams.Uptime,
			Temperature: windowsServerData.OperationalParams.Temperature,
		},
	}, nil
}
