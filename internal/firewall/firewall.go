package firewall

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// Service represents the contract for all firewall services
//
//go:generate mockgen -destination mocks/firewall.go -package mocks -source firewall.go
type Service interface {
	GetFirewallData() (MockFirewallData, error)
}

type service struct {
	config           Config
	isInitialRequest bool
}

// Config holds the firewall service related configurations, this is done to set up any required configs from settings.yaml
type Config struct {
	MockFirewallData string
}

// NewService instantiates the firewall service
func NewService(config Config) Service {
	return &service{
		config:           config,
		isInitialRequest: true,
	}
}

// GetFirewallData gives saved firewall data for first request and gives random data for subsequent requests
func (s *service) GetFirewallData() (MockFirewallData, error) {
	var mockFirewallData MockFirewallData
	if s.isInitialRequest {
		s.isInitialRequest = false
		err := json.Unmarshal([]byte(s.config.MockFirewallData), &mockFirewallData)
		if err != nil {
			logrus.Errorf("Error unmarshal firewall data:, %v ", err)
			return mockFirewallData, err
		}
		return mockFirewallData, nil
	} else {
		return MockFirewallData{
			CPUUtilization: randomFloat(1, 100),
			MemoryUsage:    randomFloat(1, 100),
			NetworkThroughput: NetworkThroughput{
				Incoming: randomFloat(500, 1500),
				Outgoing: randomFloat(800, 2000),
			},
			RequestLatencyMS:  randomFloat(1, 100),
			ErrorRate:         randomFloat(0, 10),
			PacketsFiltered:   rand.Intn(1000),
			ActiveConnections: rand.Intn(5000),
			Rules:             rand.Intn(200),
			Uptime:            randomUptime(),
			Interfaces:        generateRandomInterfaces(5),
		}, nil

	}
}

// randomFloat generates a random float number within the specified range.
func randomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano()) //TODO: Find replacement for rand.Seed
	return min + rand.Float64()*(max-min)
}

// randomString generates a random string of specified length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano()) //TODO: Find replacement for rand.Seed
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// randomUptime generates a random uptime string.
func randomUptime() string {
	days := rand.Intn(30)
	hours := rand.Intn(24)
	return fmt.Sprintf("%d days %d hours", days, hours)
}

// generateRandomInterfaces generates random data for network interfaces as a map.
func generateRandomInterfaces(count int) map[string]InterfaceStats {
	interfaces := make(map[string]InterfaceStats)
	for i := 0; i < count; i++ {
		ipAddress := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
		subnetMask := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
		packetsIn := rand.Intn(1000)
		packetsOut := rand.Intn(1000)
		interfaceName := fmt.Sprintf("eth%d", i+1)
		interfaces[interfaceName] = InterfaceStats{
			IPAddress:  ipAddress,
			SubnetMask: subnetMask,
			PacketsIn:  packetsIn,
			PacketsOut: packetsOut,
		}
	}
	return interfaces
}
