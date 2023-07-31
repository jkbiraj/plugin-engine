package plugin

import (
	"pluggin-engine/internal/firewall"
	"pluggin-engine/internal/linuxserver"
	"pluggin-engine/internal/loadbalancer"
	"pluggin-engine/internal/router"
	_switch "pluggin-engine/internal/switch"
	"pluggin-engine/internal/windowsserver"
)

// Config represents the plugin engine app related configurations
type Config struct {
	FireWallConfig      firewall.Config
	LinuxServerConfig   linuxserver.Config
	LoadBalancerConfig  loadbalancer.Config
	RouterConfig        router.Config
	SwitchConfig        _switch.Config
	WindowsServerConfig windowsserver.Config
}
