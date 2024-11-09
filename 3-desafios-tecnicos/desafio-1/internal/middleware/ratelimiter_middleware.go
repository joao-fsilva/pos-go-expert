package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"ratelimiter/internal/entity"
	"ratelimiter/internal/usecase"
)

type RateLimiterMiddleware struct {
	app *usecase.RateLimiter
}

func NewRateLimiterMiddleware(app *usecase.RateLimiter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{app: app}
}

func (m *RateLimiterMiddleware) Handler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("API_KEY")
			ip := strings.Split(r.RemoteAddr, ":")[0]

			id := ip
			rate, err := strconv.Atoi(os.Getenv("RATE"))
			if err != nil {
				http.Error(w, "Invalid rate value", http.StatusBadRequest)
				return
			}

			blockDuration, err := strconv.Atoi(os.Getenv("BLOCK_DURATION"))
			if err != nil {
				http.Error(w, "Invalid block duration value", http.StatusBadRequest)
				return
			}

			if apiKey != "" {
				id = apiKey
				newRate, newBlockDuration, err := parseApiKey(apiKey)
				if err != nil {
					http.Error(w, "Invalid API key format", http.StatusBadRequest)
					return
				}

				rate = newRate
				blockDuration = newBlockDuration
			}

			dto := usecase.RateLimiterDto{
				Id:            id,
				Rate:          rate,
				BlockDuration: time.Duration(blockDuration) * time.Minute,
			}

			log.Print("-------------------------------------------------------------------------------------------")
			output, err := m.app.Execute(dto)
			log.Print("-------------------------------------------------------------------------------------------")

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if output.Error != nil {
				var limiterErr *entity.LimiterError
				if errors.As(output.Error, &limiterErr) {
					switch limiterErr.Err {
					case "entity_not_found":
						break
					case "is_blocked":
						http.Error(w, limiterErr.Error(), http.StatusTooManyRequests)
						return
					case "expired_limiter":
						break
					default:
						http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
						return
					}
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func parseApiKey(apiKey string) (int, int, error) {
	parts := strings.Split(apiKey, "_")
	if len(parts) < 5 {
		return 0, 0, fmt.Errorf("invalid API key format")
	}

	rate, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid rate value: %v", err)
	}

	blockDuration, err := strconv.Atoi(parts[4])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid block duration value: %v", err)
	}

	return rate, blockDuration, nil
}
