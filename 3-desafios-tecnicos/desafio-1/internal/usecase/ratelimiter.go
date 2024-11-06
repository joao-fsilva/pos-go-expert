package usecase

import (
	"errors"
	"ratelimiter/internal/entity"
	"time"
)

type RateLimiterDto struct {
	Id            string
	Rate          int
	BlockDuration time.Duration
}

type RateLimiterOutputDto struct {
	Error   error
	Success bool
}

type RateLimiter struct {
	limiterRepository entity.LimiterRepository
}

func NewRateLimiter(limiterRepository entity.LimiterRepository) *RateLimiter {
	return &RateLimiter{
		limiterRepository: limiterRepository,
	}
}

func (rl *RateLimiter) Execute(dto RateLimiterDto) (RateLimiterOutputDto, error) {
	e, err := rl.limiterRepository.Find(dto.Id)
	var errOutput error

	if err != nil {
		var limiterErr *entity.LimiterError
		if errors.As(err, &limiterErr) && limiterErr.Err == "entity_not_found" {
			e = entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
			errOutput = e.IncrementAccessCount()

			if err = rl.limiterRepository.Create(e); err != nil {
				return RateLimiterOutputDto{}, err
			}

			return RateLimiterOutputDto{Error: errOutput, Success: true}, nil
		}

		return RateLimiterOutputDto{}, err
	}

	errOutput = e.IncrementAccessCount()
	if errOutput != nil {
		var limiterErr *entity.LimiterError
		if errors.As(errOutput, &limiterErr) && limiterErr.Err == "expired_limiter" {
			e = entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
			errOutput = e.IncrementAccessCount()
			if err = rl.limiterRepository.Update(e); err != nil {
				return RateLimiterOutputDto{}, err
			}
			return RateLimiterOutputDto{Error: nil, Success: true}, nil
		}

		if err = rl.limiterRepository.Update(e); err != nil {
			return RateLimiterOutputDto{}, err
		}

		return RateLimiterOutputDto{Error: errOutput, Success: false}, nil
	}

	if err = rl.limiterRepository.Update(e); err != nil {
		return RateLimiterOutputDto{}, err
	}

	return RateLimiterOutputDto{Error: nil, Success: true}, nil
}
