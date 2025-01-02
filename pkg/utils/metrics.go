package utils

import (
	"fmt"
	"sync"
	"time"
)

var (
	startTime     = time.Now()
	totalRequests int
	totalErrors   int
	mu            sync.Mutex
)

// IncrementRequest increments the total request count
func IncrementRequest() {
	mu.Lock()
	defer mu.Unlock()
	totalRequests++
}

// IncrementError increments the total error count
func IncrementError() {
	mu.Lock()
	defer mu.Unlock()
	totalErrors++
}

// GetMetrics returns the calculated metrics
func GetMetrics() (uptime string, requestCount string, errorRate string) {
	mu.Lock()
	defer mu.Unlock()

	// Calculate uptime
	uptimeDuration := time.Since(startTime)
	uptime = uptimeDuration.String()

	// Format request count
	requestCount = formatCount(totalRequests)

	// Calculate error rate
	var rate float64
	if totalRequests > 0 {
		rate = (float64(totalErrors) / float64(totalRequests)) * 100
	}
	errorRate = formatRate(rate)

	return uptime, requestCount, errorRate
}

// formatCount formats an integer as a string
func formatCount(count int) string {
	return time.Duration(count).String()
}

// formatRate formats a float as a percentage string
func formatRate(rate float64) string {
	return fmt.Sprintf("%.2f%%", rate)
}
