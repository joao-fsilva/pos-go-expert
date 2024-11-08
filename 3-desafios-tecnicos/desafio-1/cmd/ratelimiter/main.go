package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"ratelimiter/internal/entity"
	"ratelimiter/internal/infra"
	"ratelimiter/internal/usecase"
	"time"
)

func main() {
	r := chi.NewRouter()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	repo := infra.NewLimiterRepositoryRedis(rdb)

	app := usecase.NewRateLimiter(repo)

	r.Use(RateLimiterMiddleware(app))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}

func RateLimiterMiddleware(app *usecase.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			dto := usecase.RateLimiterDto{
				Id:            ip,
				Rate:          5,
				BlockDuration: 1 * time.Minute,
			}

			log.Print("-------------------------------------------------------------------------------------------")
			output, err := app.Execute(dto)
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
