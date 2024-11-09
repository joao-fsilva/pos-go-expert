package usecase

import (
	"errors"
	"log"
	"ratelimiter/internal/entity"
	"sync"
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
	isProcessingMap   map[string]*sync.Mutex
	isProcessingMutex *sync.Mutex
}

func NewRateLimiter(limiterRepository entity.LimiterRepository) *RateLimiter {
	return &RateLimiter{
		limiterRepository: limiterRepository,
		isProcessingMap:   make(map[string]*sync.Mutex),
		isProcessingMutex: &sync.Mutex{},
	}
}

func (rl *RateLimiter) Execute(dto RateLimiterDto) (RateLimiterOutputDto, error) {
	log.Printf("Executing rate limiter for ID: %s, Rate: %d, Block Duration: %s", dto.Id, dto.Rate, dto.BlockDuration)

	rl.isProcessingMutex.Lock()
	mutex, exists := rl.isProcessingMap[dto.Id]
	if !exists {
		mutex = &sync.Mutex{}
		rl.isProcessingMap[dto.Id] = mutex
	}
	rl.isProcessingMutex.Unlock()

	mutex.Lock()
	defer mutex.Unlock()

	e, err := rl.limiterRepository.Find(dto.Id)
	if err != nil {
		var limiterErr *entity.LimiterError
		if errors.As(err, &limiterErr) && limiterErr.Err == "entity_not_found" {
			log.Println("Limiter entity not found, creating a new one.")
			e = entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
			errOutput := e.IncrementAccessCount()
			log.Print("Incremented access count for new limiter")

			if err = rl.limiterRepository.Create(e); err != nil {
				log.Printf("Error creating limiter for ID: %s, error: %v", dto.Id, err)
				return RateLimiterOutputDto{}, err
			}

			log.Printf("Limiter created successfully for ID: %s", dto.Id)
			return RateLimiterOutputDto{Error: errOutput, Success: true}, nil
		}

		return RateLimiterOutputDto{}, err
	}

	log.Printf("Limiter found for ID: %s, incrementing access count.", dto.Id)
	errOutput := e.IncrementAccessCount()
	if errOutput != nil {
		log.Printf("Error incrementing access count for limiter ID: %s, error: %v", dto.Id, errOutput)
		var limiterErr *entity.LimiterError
		if errors.As(errOutput, &limiterErr) && limiterErr.Err == "expired_limiter" {
			log.Println("Limiter has expired, creating a new one.")
			e = entity.NewLimiter(dto.Id, dto.Rate, dto.BlockDuration)
			errOutput = e.IncrementAccessCount()
			if err = rl.limiterRepository.Update(e); err != nil {
				log.Printf("Error updating limiter for ID: %s, error: %v", dto.Id, err)
				return RateLimiterOutputDto{}, err
			}
			log.Printf("Limiter updated successfully after expiration for ID: %s", dto.Id)
			return RateLimiterOutputDto{Error: nil, Success: true}, nil
		}

		if err = rl.limiterRepository.Update(e); err != nil {
			log.Printf("Error updating limiter for ID: %s, error: %v", dto.Id, err)
			return RateLimiterOutputDto{}, err
		}

		return RateLimiterOutputDto{Error: errOutput, Success: false}, nil
	}

	if err = rl.limiterRepository.Update(e); err != nil {
		log.Printf("Error updating limiter for ID: %s, error: %v", dto.Id, err)
		return RateLimiterOutputDto{}, err
	}

	log.Printf("Limiter updated successfully for ID: %s", dto.Id)
	return RateLimiterOutputDto{Error: nil, Success: true}, nil
}
