package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"ratelimiter/internal/ratelimiter"
)

func main() {
	r := chi.NewRouter()

	rl := ratelimiter.NewRateLimiter(2)

	r.Use(RateLimiterMiddleware(rl))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}

func RateLimiterMiddleware(rl *ratelimiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr // Obtém o IP do cliente

			if !rl.IsAllowed(ip) {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
