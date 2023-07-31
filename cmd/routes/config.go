package routes

import (
	"pluggin-engine/internal/plugin"
)

// Config contains everything needed to create new routes
type Config struct {
	Port                    string
	queryParamVirtualDevice string
	PluginService           plugin.Service
}

type handler struct {
	Config
}

const queryParamVirtualDevice = "virtualDevice"

func newHandler(config Config) handler {
	if config.PluginService == nil {
		return handler{}
	}
	config.queryParamVirtualDevice = queryParamVirtualDevice
	return handler{
		config,
	}
}
