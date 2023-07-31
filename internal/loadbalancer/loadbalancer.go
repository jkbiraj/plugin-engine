package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// Service represents the contract for all firewall services
//
//go:generate mockgen -destination mocks/loadbalancer.go -package mocks -source loadbalancer.go
type Service interface {
	GetLoadBalancerData() (MockLoadBalancerData, error)
}

type service struct {
	config           Config
	isInitialRequest bool
}

// Config holds the load balancer service related configurations, this is done to set up any required configs from settings.yaml
type Config struct {
	MockLoadBalancerData string
}

// NewService instantiates the load balancer service
func NewService(config Config) Service {
	return &service{
		config:           config,
		isInitialRequest: true,
	}
}

// GetLoadBalancerData gives saved load balancer data for first request and gives random data for subsequent requests
func (s *service) GetLoadBalancerData() (MockLoadBalancerData, error) {
	var mockLoadBalancerData MockLoadBalancerData
	if s.isInitialRequest {
		s.isInitialRequest = false
		err := json.Unmarshal([]byte(s.config.MockLoadBalancerData), &mockLoadBalancerData)
		if err != nil {
			logrus.Errorf("Error unmarshal load balancer data:, %v ", err)
			return mockLoadBalancerData, err
		}
		return mockLoadBalancerData, nil
	} else {

		return MockLoadBalancerData{
			CPUUtilization:    rand.Float64() * 100,
			MemoryUsage:       rand.Float64() * 100,
			NetworkThroughput: NetworkThroughput{Incoming: rand.Float64() * 1000, Outgoing: rand.Float64() * 1000},
			RequestLatencyMS:  rand.Float64() * 50,
			ErrorRate:         rand.Float64() * 0.1,
			ActiveConnections: rand.Intn(5000),
			Uptime:            randomUptime(),
		}, nil
	}
}

// randomUptime generates a random uptime in the format "X days Y hours"
func randomUptime() string {
	days := rand.Intn(365)
	hours := rand.Intn(24)
	return fmt.Sprintf("%d days %d hours", days, hours)
}
