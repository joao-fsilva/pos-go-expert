package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"ratelimiter/internal/infra"
	"ratelimiter/internal/middleware"
	"ratelimiter/internal/usecase"
)

func main() {
	if err := godotenv.Load("cmd/ratelimiter/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	r := chi.NewRouter()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	repo := infra.NewLimiterRepositoryRedis(rdb)
	app := usecase.NewRateLimiter(repo)

	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(app)

	r.Use(rateLimiterMiddleware.Handler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
