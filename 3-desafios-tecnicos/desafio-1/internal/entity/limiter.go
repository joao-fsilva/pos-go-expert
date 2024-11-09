package entity

import "time"

type Limiter struct {
	Id            string
	Rate          int
	AccessCount   int
	BlockedAt     time.Time
	BlockDuration time.Duration
	CreatedAt     time.Time
}

func NewLimiter(id string, rate int, blockDuration time.Duration) *Limiter {
	return &Limiter{
		Id:            id,
		Rate:          rate,
		AccessCount:   0,
		BlockedAt:     time.Time{},
		BlockDuration: blockDuration,
		CreatedAt:     time.Now(),
	}
}

func (l *Limiter) IncrementAccessCount() error {
	if l.BlockedAt.IsZero() == false {
		if time.Since(l.BlockedAt) < l.BlockDuration {
			return NewIncrementBlockedError()
		}

		return NewExpiredLimiterError()
	}

	if time.Since(l.CreatedAt) > time.Second {
		return NewExpiredLimiterError()
	}

	if (l.AccessCount + 1) > l.Rate {
		l.BlockedAt = time.Now()
		return NewIncrementBlockedError()
	}

	l.AccessCount++

	return nil
}
