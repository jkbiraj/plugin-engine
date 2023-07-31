package main

import (
	"log"
	"time"

	"pluggin-engine/internal/firewall"
	"pluggin-engine/internal/linuxserver"
	"pluggin-engine/internal/loadbalancer"
	"pluggin-engine/internal/plugin"
	"pluggin-engine/internal/router"
	_switch "pluggin-engine/internal/switch"
	"pluggin-engine/internal/windowsserver"

	"pluggin-engine/cmd/config"
	"pluggin-engine/cmd/routes"
)

func startPluginEngine() error {
	startTime := time.Now()
	if err := config.Set(); err != nil {
		return err
	}
	pluginService := NewPluginService()
	routesConfig := config.Routes(pluginService)
	err := routes.SetAndServeRouter(routesConfig)
	if err != nil {
		return err
	}
	log.Printf("Application started on port 8443, set up took %vms", time.Since(startTime)) //TODO: This debug line doesn't work currently, Look for alternate solution
	return nil
}

// NewPluginService creates an instance of the main plugin service
func NewPluginService() plugin.Service {
	pluginConfig := config.NewPluginConfig()
	firewallService := firewall.NewService(pluginConfig.FireWallConfig)
	linuxServerService := linuxserver.NewService(pluginConfig.LinuxServerConfig)
	loadBalancerService := loadbalancer.NewService(pluginConfig.LoadBalancerConfig)
	routerService := router.NewService(pluginConfig.RouterConfig)
	switchService := _switch.NewService(pluginConfig.SwitchConfig)
	windowsServerService := windowsserver.NewService(pluginConfig.WindowsServerConfig)

	return plugin.NewService(pluginConfig, firewallService, linuxServerService, loadBalancerService, routerService, switchService,
		windowsServerService)
}
