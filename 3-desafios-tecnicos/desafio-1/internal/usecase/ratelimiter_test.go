package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"ratelimiter/internal/entity"
	"ratelimiter/internal/infra"
	"testing"
	"time"
)

func setupRepo() entity.LimiterRepository {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	client.FlushDB(context.Background())

	return infra.NewLimiterRepositoryRedis(client)
}

func TestExecuteWhenLimiterNotFoundThenCreatesLimiter(t *testing.T) {
	repo := setupRepo()
	usecase := NewRateLimiter(repo)

	dto := RateLimiterDto{
		Id:            "test_id",
		Rate:          5,
		BlockDuration: 5 * time.Minute,
	}

	output, err := usecase.Execute(dto)
	assert.NoError(t, err)
	assert.True(t, output.Success)
	assert.Nil(t, output.Error)

	limiter, err := repo.Find("test_id")
	assert.NoError(t, err)
	assert.Equal(t, 1, limiter.AccessCount)
}

func TestExecuteWhenLimiterExistsAndExpiredThenResetsLimiter(t *testing.T) {
	repo := setupRepo()
	usecase := NewRateLimiter(repo)

	dto := RateLimiterDto{
		Id:            "test_id",
		Rate:          1,
		BlockDuration: 1 * time.Second,
	}

	limiter := entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
	limiter.CreatedAt = time.Now().Add(-2 * time.Second)
	repo.Create(limiter)

	output, err := usecase.Execute(dto)
	assert.NoError(t, err)
	assert.True(t, output.Success)
	assert.Nil(t, output.Error)

	limiter, err = repo.Find("test_id")
	assert.NoError(t, err)
	assert.Equal(t, 1, limiter.AccessCount)
}

func TestExecuteWhenLimiterExistsAndNotExpiredThenIncrementsAccessCount(t *testing.T) {
	repo := setupRepo()
	usecase := NewRateLimiter(repo)

	dto := RateLimiterDto{
		Id:            "test_id",
		Rate:          5,
		BlockDuration: 5 * time.Minute,
	}

	limiter := entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
	repo.Create(limiter)

	output, err := usecase.Execute(dto)
	assert.NoError(t, err)
	assert.True(t, output.Success)
	assert.Nil(t, output.Error)

	limiter, err = repo.Find("test_id")
	assert.NoError(t, err)
	assert.Equal(t, 1, limiter.AccessCount)
}

func TestExecuteWhenLimiterBlockedThenReturnsError(t *testing.T) {
	repo := setupRepo()
	usecase := NewRateLimiter(repo)

	dto := RateLimiterDto{
		Id:            "test_id",
		Rate:          2,
		BlockDuration: 5 * time.Minute,
	}

	limiter := entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
	limiter.IncrementAccessCount()
	limiter.IncrementAccessCount()
	repo.Create(limiter)

	output, err := usecase.Execute(dto)
	assert.NoError(t, err)
	assert.False(t, output.Success)
	assert.Equal(t, "is_blocked", output.Error.(*entity.LimiterError).Err)
}
