// internal/ratelimiter/ratelimiter.go
package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate  int
	times map[string][]time.Time
	mu    sync.Mutex
}

func NewRateLimiter(rate int) *RateLimiter {
	return &RateLimiter{
		rate:  rate,
		times: make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) IsAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if _, exists := rl.times[ip]; !exists {
		rl.times[ip] = []time.Time{}
	}

	threshold := now.Add(-time.Second)
	var newTimes []time.Time
	for _, t := range rl.times[ip] {
		if t.After(threshold) {
			newTimes = append(newTimes, t)
		}
	}

	rl.times[ip] = newTimes

	if len(rl.times[ip]) < rl.rate {
		rl.times[ip] = append(rl.times[ip], now)
		return true
	}

	return false
}
