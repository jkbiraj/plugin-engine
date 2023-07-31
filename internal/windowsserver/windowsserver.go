package windowsserver

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// Service represents the contract for all firewall services
//
//go:generate mockgen -destination mocks/windowsserver.go -package mocks -source windowsserver.go
type Service interface {
	GetWindowsServerData() (MockWindowsServerData, error)
}

type service struct {
	config           Config
	isInitialRequest bool
}

// Config holds the Windows server service related configurations, this is done to set up any required configs from settings.yaml
type Config struct {
	MockWindowsServerData string
}

// NewService instantiates the Windows server service
func NewService(config Config) Service {
	return &service{
		config:           config,
		isInitialRequest: true,
	}
}

// GetWindowsServerData gives saved Windows server data for first request and gives random data for subsequent requests
func (s *service) GetWindowsServerData() (MockWindowsServerData, error) {
	var mockWindowsServerData MockWindowsServerData
	if s.isInitialRequest {
		s.isInitialRequest = false
		err := json.Unmarshal([]byte(s.config.MockWindowsServerData), &mockWindowsServerData)
		if err != nil {
			logrus.Errorf("Error unmarshal windows server data:, %v ", err)
			return mockWindowsServerData, err
		}
		return mockWindowsServerData, nil
	} else {
		return MockWindowsServerData{
			ServerID: "server-" + randomString(4),
			CPUUtilization: CPUUtilization{
				Average: randomFloat(10, 80),
				Max:     randomFloat(30, 90),
				Min:     randomFloat(5, 20),
			},
			MemoryUsage: MemoryUsage{
				Used:       rand.Intn(4096) + 1024,
				Total:      8192,
				Percentage: randomFloat(30, 70),
			},
			NetworkThroughput: NetworkThroughput{
				Incoming: randomFloat(500, 1500),
				Outgoing: randomFloat(800, 2000),
			},
			RequestLatency: RequestLatency{
				AverageMs: randomFloat(5, 15),
				MaxMs:     randomFloat(10, 25),
			},
			ErrorRates: ErrorRates{
				HTTP500: randomFloat(0.001, 0.1),
				HTTP404: randomFloat(0.001, 0.05),
			},
			OperationalParams: OperationalParams{
				IsOnline:    rand.Intn(2) == 1,
				Uptime:      randomUptime(),
				Temperature: randomFloat(30, 40),
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
