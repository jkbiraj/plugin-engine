package linuxserver

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// Service represents the contract for all firewall services
//
//go:generate mockgen -destination mocks/linuxserver.go -package mocks -source linuxserver.go
type Service interface {
	GetLinuxServerData() (MockLinuxServerData, error)
}

type service struct {
	config           Config
	isInitialRequest bool
}

// Config holds the linux server service related configurations, this is done to set up any required configs from settings.yaml
type Config struct {
	MockLinuxServerData string
}

// NewService instantiates the linux server service
func NewService(config Config) Service {
	return &service{
		config:           config,
		isInitialRequest: true,
	}
}

// GetLinuxServerData gives saved linux server data for first request and gives random data for subsequent requests
func (s *service) GetLinuxServerData() (MockLinuxServerData, error) {
	var mockLinuxServerData MockLinuxServerData
	if s.isInitialRequest {
		s.isInitialRequest = false
		err := json.Unmarshal([]byte(s.config.MockLinuxServerData), &mockLinuxServerData)
		if err != nil {
			logrus.Errorf("Error unmarshal linux server data:, %v ", err)
			return mockLinuxServerData, err
		}
		return mockLinuxServerData, nil
	} else {
		return MockLinuxServerData{
			CPUUtilization: randomFloat(1, 100),
			MemoryUsage:    randomFloat(1, 100),
			NetworkThroughput: NetworkThroughput{
				Sent:     rand.Intn(5000) + 1000, // Random sent data between 1000 and 6000
				Received: rand.Intn(5000) + 1000, // Random received data between 1000 and 6000
			},
			RequestLatencyMS: randomFloat(1, 100),
			ErrorRate:        randomFloat(0, 10),
			DiskUsage: DiskUsage{
				Total:     rand.Uint64()%(500*1024*1024*1024) + 100*1024*1024*1024, // Random total disk space between 100GB and 600GB
				Used:      rand.Uint64()%(400*1024*1024*1024) + 50*1024*1024*1024,  // Random used disk space between 50GB and 450GB
				Available: 0,
			},

			ActiveConnections: rand.Intn(1000),
			Processes:         rand.Intn(500),
			Uptime:            randomUptime(),
			LoadAverage: LoadAverage{
				Min1:  randomFloat(0, 10),
				Min5:  randomFloat(0, 10),
				Min15: randomFloat(0, 10),
			},
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
