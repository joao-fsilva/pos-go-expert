package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"ratelimiter/internal/entity"
)

type LimiterRepositoryRedis struct {
	client *redis.Client
	ctx    context.Context
}

func NewLimiterRepositoryRedis(redisClient *redis.Client) *LimiterRepositoryRedis {
	return &LimiterRepositoryRedis{
		client: redisClient,
		ctx:    context.Background(),
	}
}

func (r *LimiterRepositoryRedis) Create(limiter *entity.Limiter) error {
	data, err := json.Marshal(limiter)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, limiter.Id, data, 0).Err()
}

func (r *LimiterRepositoryRedis) Update(limiter *entity.Limiter) error {
	err := r.Create(limiter)
	if err != nil {
		return err
	}
	return nil
}

func (r *LimiterRepositoryRedis) Find(id string) (*entity.Limiter, error) {
	data, err := r.client.Get(r.ctx, id).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var limiter entity.Limiter
	err = json.Unmarshal([]byte(data), &limiter)
	if err != nil {
		return nil, err
	}

	return &limiter, nil
}
