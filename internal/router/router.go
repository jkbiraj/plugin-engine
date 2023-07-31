package router

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// Service represents the contract for all firewall services
//
//go:generate mockgen -destination mocks/router.go -package mocks -source router.go
type Service interface {
	GetRouterData() (MockRouterData, error)
}

type service struct {
	config           Config
	isInitialRequest bool
}

// Config holds the router service related configurations, this is done to set up any required configs from settings.yaml
type Config struct {
	MockRouterData string
}

// NewService instantiates the router service
func NewService(config Config) Service {
	return &service{
		config:           config,
		isInitialRequest: true,
	}
}

// GetRouterData gives saved router data for first request and gives random data for subsequent requests
func (s *service) GetRouterData() (MockRouterData, error) {
	var mockRouterData MockRouterData
	if s.isInitialRequest {
		s.isInitialRequest = false
		err := json.Unmarshal([]byte(s.config.MockRouterData), &mockRouterData)
		if err != nil {
			logrus.Errorf("Error unmarshal router data:, %v ", err)
			return mockRouterData, err
		}
		return mockRouterData, nil
	} else {
		return MockRouterData{
			RouterID: "router-" + randomString(4),
			CPUUtilization: CPUUtilization{
				Average: randomFloat(10, 80),
				Max:     randomFloat(50, 100),
				Min:     randomFloat(5, 20),
			},
			MemoryUsage: MemoryUsage{Used: rand.Intn(2048) + 1024,
				Total:      4096,
				Percentage: randomFloat(30, 70),
			},
			NetworkThroughput: NetworkThroughput{Incoming: randomFloat(100, 1000),
				Outgoing: randomFloat(200, 1200),
			},
			RequestLatency: RequestLatency{
				AverageMs: randomFloat(5, 30),
				MaxMs:     randomFloat(20, 50),
			},
			ErrorRates: ErrorRates{
				HTTP500: randomFloat(0, 1),
				HTTP404: randomFloat(0, 0.5),
			},
			OperationalParams: OperationalParams{IsActive: rand.Intn(2) == 1,
				Uptime:      randomUptime(),
				Temperature: randomFloat(30, 50),
			},
		}, nil
	}

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

// randomFloat generates a random float number within the specified range.
func randomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano()) //TODO: Find replacement for rand.Seed
	return min + rand.Float64()*(max-min)
}

// randomUptime generates a random uptime string.
func randomUptime() string {
	days := rand.Intn(30)
	hours := rand.Intn(24)
	return fmt.Sprintf("%d days %d hours", days, hours)
}
