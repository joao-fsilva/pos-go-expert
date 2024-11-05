package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLimiterWhenCalledThenSuccess(t *testing.T) {
	limiter := NewLimiter("test_id", 5, 5*time.Minute)

	assert.Equal(t, "test_id", limiter.Id)
	assert.Equal(t, 5, limiter.Rate)
	assert.Equal(t, 0, limiter.AccessCount)
	assert.True(t, limiter.BlockedAt.IsZero())
	assert.Equal(t, 5*time.Minute, limiter.BlockDuration)
	assert.WithinDuration(t, time.Now(), limiter.CreatedAt, time.Second)
}

func TestIncrementAccessCountWhenItDidNotExceedTheRateThenSuccess(t *testing.T) {
	limiter := NewLimiter("test_id", 5, 5*time.Minute)

	for i := 0; i < 5; i++ {
		err := limiter.IncrementAccessCount()
		assert.NoError(t, err)
		assert.Equal(t, i+1, limiter.AccessCount)
	}
}

func TestIncrementAccessCountWhenAccessExceededTheRateThenBlockAndIsBlockedError(t *testing.T) {
	limiter := NewLimiter("test_id", 5, 5*time.Minute)

	for i := 0; i < 5; i++ {
		_ = limiter.IncrementAccessCount()
	}

	err := limiter.IncrementAccessCount()
	assert.Error(t, err)
	assert.Equal(t, "is_blocked", err.(*LimiterError).Err)
	assert.Equal(t, 6, limiter.AccessCount)
	assert.False(t, limiter.BlockedAt.IsZero())
}

func TestIncrementAccessCountWhenBlockExpiresThenExpiredLimitError(t *testing.T) {
	limiter := NewLimiter("test_id", 1, 1*time.Second)

	err := limiter.IncrementAccessCount()
	assert.NoError(t, err)
	assert.Equal(t, 1, limiter.AccessCount)

	_ = limiter.IncrementAccessCount()
	time.Sleep(2 * time.Second)
	err = limiter.IncrementAccessCount()
	assert.Equal(t, "expired_limiter", err.(*LimiterError).Err)
}

func TestIncrementAccessCountWhenLimitExpiresThenExpiredLimitError(t *testing.T) {
	limiter := NewLimiter("test_id", 1, 5*time.Minute)

	limiter.CreatedAt = time.Now().Add(-2 * time.Second)

	err := limiter.IncrementAccessCount()
	assert.Error(t, err)
	assert.Equal(t, "expired_limiter", err.(*LimiterError).Err)
}
