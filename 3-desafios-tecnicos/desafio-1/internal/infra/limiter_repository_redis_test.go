package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"ratelimiter/internal/entity"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	repo := NewLimiterRepositoryRedis(rdb)

	limiter := entity.NewLimiter("test_id", 5, 5)

	err := repo.Create(limiter)

	assert.NoError(t, err)

	result, err := repo.Find(limiter.Id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, limiter.Id, result.Id)
	assert.Equal(t, limiter.Rate, result.Rate)
	assert.Equal(t, limiter.AccessCount, result.AccessCount)
	assert.Equal(t, limiter.BlockedAt, result.BlockedAt)
	assert.Equal(t, limiter.BlockDuration, result.BlockDuration)
	assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)
}

func TestUpdate(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	repo := NewLimiterRepositoryRedis(rdb)

	limiter := entity.NewLimiter("test_id", 5, 5)

	err := repo.Create(limiter)

	err2 := limiter.IncrementAccessCount()

	err3 := repo.Update(limiter)

	assert.NoError(t, err)
	assert.NoError(t, err2)
	assert.NoError(t, err3)

	result, err := repo.Find(limiter.Id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, limiter.Id, result.Id)
	assert.Equal(t, limiter.Rate, result.Rate)
	assert.Equal(t, limiter.AccessCount, result.AccessCount)
	assert.Equal(t, limiter.BlockedAt, result.BlockedAt)
	assert.Equal(t, limiter.BlockDuration, result.BlockDuration)
	assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)
}
