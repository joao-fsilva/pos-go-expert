package entity

import "time"

type Limiter struct {
	id            string
	rate          int
	accessCount   int
	blockedAt     time.Time
	blockDuration time.Duration
	createdAt     time.Time
}

func NewLimiter(id string, rate int, blockDuration time.Duration) *Limiter {
	return &Limiter{
		id:            id,
		rate:          rate,
		accessCount:   0,
		blockedAt:     time.Time{},
		blockDuration: blockDuration,
		createdAt:     time.Now(),
	}
}

func Restore(id string, rate int, accessCount int, blockedAt time.Time, blockDuration time.Duration, createdAt time.Time) *Limiter {
	return &Limiter{
		id:            id,
		rate:          rate,
		accessCount:   accessCount,
		blockedAt:     blockedAt,
		blockDuration: blockDuration,
		createdAt:     createdAt,
	}
}

func (l *Limiter) IncrementAccessCount() error {
	if l.blockedAt.IsZero() == false {
		if time.Since(l.blockedAt) < l.blockDuration {
			return NewIncrementBlockedError()
		}

		return NewExpiredLimiterError()
	}

	if time.Since(l.createdAt) > time.Second {
		return NewExpiredLimiterError()
	}

	l.accessCount++
	if l.accessCount > l.rate {
		l.blockedAt = time.Now()
		return NewIncrementBlockedError()
	}

	return nil
}
