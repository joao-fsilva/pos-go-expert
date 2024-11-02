package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"weather-zip-code/controller"
	"weather-zip-code/internal/infra"
	"weather-zip-code/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}

	httpTracer := infra.NewHttpTracer("serviceb")

	zipCodeService := infra.NewZipCodeServiceViaCep(httpTracer)
	weatherService := infra.NewWeatherServiceWeatherApi(apiKey)

	app := usecase.NewGetWeatherByZipCode(zipCodeService, weatherService)

	weatherController := controller.NewWeatherController(app)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/weather", weatherController.Handle)

	http.ListenAndServe(":8080", nil)
}
