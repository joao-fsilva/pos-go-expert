package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		time.Sleep(5 * time.Second)

		w.Write([]byte("Hello, World!"))
	})

	go func() {
		fmt.Println("Server running on port 3000")

		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on port 3000: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1) //sinal de interrupção pelo sistema operacional
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}

	fmt.Println("Server stopped")
}
