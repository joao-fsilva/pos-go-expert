package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				debug.PrintStack()
				http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("panic")
	})

	log.Println("Listening on :3000")

	if err := http.ListenAndServe(":3000", recoveryMiddleware(mux)); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
