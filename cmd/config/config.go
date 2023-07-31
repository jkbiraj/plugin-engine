package config

import (
	"log"

	"pluggin-engine/cmd/routes"
	"pluggin-engine/internal/firewall"
	"pluggin-engine/internal/linuxserver"
	"pluggin-engine/internal/loadbalancer"
	"pluggin-engine/internal/plugin"
	"pluggin-engine/internal/router"
	_switch "pluggin-engine/internal/switch"
	"pluggin-engine/internal/windowsserver"

	"github.com/spf13/viper"
)

// Registry is for the configuration values.
var Registry *viper.Viper

// Set the configs
func Set() error {

	Registry = viper.New()

	Registry.SetConfigName("settings") // name of the config file, avoid any extensions
	Registry.SetConfigType("yaml")     // or json, toml etc

	// Add paths to search for the config file
	Registry.AddConfigPath(".")
	Registry.AddConfigPath("./settings")

	// Read configuration file
	if err := Registry.ReadInConfig(); err != nil {
		log.Printf("Error reading settings file: %v", err)
		return err
	}
	return nil
}

// Routes returns the configuration required for the routes/handlers (REST interface)
func Routes(pluginService plugin.Service) routes.Config {
	return routes.Config{
		Port:          Registry.GetString("SERVER_PORT"),
		PluginService: pluginService,
	}
}

// NewPluginConfig creates a new config for the service
func NewPluginConfig() plugin.Config {
	return plugin.Config{
		FireWallConfig: firewall.Config{
			MockFirewallData: Registry.GetString("MOCK_FIREWALL_DATA"),
		},
		LinuxServerConfig: linuxserver.Config{
			MockLinuxServerData: Registry.GetString("MOCK_LINUX_SERVER_DATA"),
		},
		LoadBalancerConfig: loadbalancer.Config{
			MockLoadBalancerData: Registry.GetString("MOCK_LOAD_BALANCER_DATA"),
		},
		RouterConfig: router.Config{
			MockRouterData: Registry.GetString("MOCK_ROUTER_DATA"),
		},
		SwitchConfig: _switch.Config{
			MockSwitchData: Registry.GetString("MOCK_SWITCH_DATA"),
		},
		WindowsServerConfig: windowsserver.Config{
			MockWindowsServerData: Registry.GetString("MOCK_WINDOWS_SERVER_DATA"),
		},
	}
}
