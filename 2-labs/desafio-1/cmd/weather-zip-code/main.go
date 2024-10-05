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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	zipCodeService := infra.NewZipCodeServiceViaCep()
	weatherService := infra.NewWeatherServiceWeatherApi(apiKey)

	app := usecase.NewGetWeatherByZipCode(zipCodeService, weatherService)

	weatherController := controller.NewWeatherController(app)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/weather", weatherController.Handle)

	http.ListenAndServe(":8080", nil)
}
