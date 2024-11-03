package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewLimiterWhenCalledThenSuccess(t *testing.T) {
	limiter := NewLimiter("test_id", 5, 5*time.Minute)

	assert.Equal(t, "test_id", limiter.id)
	assert.Equal(t, 5, limiter.rate)
	assert.Equal(t, 0, limiter.accessCount)
	assert.True(t, limiter.blockedAt.IsZero())
	assert.Equal(t, 5*time.Minute, limiter.blockDuration)
	assert.WithinDuration(t, time.Now(), limiter.createdAt, time.Second)
}

func TestRestoreWhenCalledThenSuccess(t *testing.T) {
	id := "user1"
	rate := 5
	accessCount := 2
	blockedAt := time.Now().Add(-10 * time.Minute)
	blockDuration := 5 * time.Minute
	createdAt := time.Now().Add(-1 * time.Hour)

	limiter := Restore(id, rate, accessCount, blockedAt, blockDuration, createdAt)

	assert.NotNil(t, limiter)
	assert.Equal(t, id, limiter.id)
	assert.Equal(t, rate, limiter.rate)
	assert.Equal(t, accessCount, limiter.accessCount)
	assert.Equal(t, blockedAt.Unix(), limiter.blockedAt.Unix())
	assert.Equal(t, blockDuration, limiter.blockDuration)
	assert.Equal(t, createdAt.Unix(), limiter.createdAt.Unix())
}

func TestIncrementAccessCountWhenItDidNotExceedTheRateThenSuccess(t *testing.T) {
	limiter := NewLimiter("test_id", 5, 5*time.Minute)

	for i := 0; i < 5; i++ {
		err := limiter.IncrementAccessCount()
		assert.NoError(t, err)
		assert.Equal(t, i+1, limiter.accessCount)
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
	assert.Equal(t, 6, limiter.accessCount)
	assert.False(t, limiter.blockedAt.IsZero())
}

func TestIncrementAccessCountWhenBlockExpiresThenExpiredLimitError(t *testing.T) {
	limiter := NewLimiter("test_id", 1, 1*time.Second)

	err := limiter.IncrementAccessCount()
	assert.NoError(t, err)
	assert.Equal(t, 1, limiter.accessCount)

	_ = limiter.IncrementAccessCount()
	time.Sleep(2 * time.Second)
	err = limiter.IncrementAccessCount()
	assert.Equal(t, "expired_limiter", err.(*LimiterError).Err)
}

func TestIncrementAccessCountWhenLimitExpiresThenExpiredLimitError(t *testing.T) {
	limiter := NewLimiter("test_id", 1, 5*time.Minute)

	limiter.createdAt = time.Now().Add(-2 * time.Second)

	err := limiter.IncrementAccessCount()
	assert.Error(t, err)
	assert.Equal(t, "expired_limiter", err.(*LimiterError).Err)
}
