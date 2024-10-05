package usecase

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
	"weather-zip-code/internal/entity"
	"weather-zip-code/internal/infra"
)

func LoadUseCase() *GetWeatherByZipCode {
	envFile := filepath.Join("..", "..", ".env.test")

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	zipCodeService := infra.NewZipCodeServiceViaCep()
	weatherService := infra.NewWeatherServiceWeatherApi(apiKey)

	return NewGetWeatherByZipCode(zipCodeService, weatherService)
}

func TestGivenAnInvalidZipCode_WhenGetWeather_ThenShouldReturnError(t *testing.T) {
	useCase := LoadUseCase()

	var dto GetWeatherByZipCodeDTO
	dto.ZipCode = "123456789"

	_, err := useCase.Execute(dto)

	assert.ErrorIs(t, err, entity.ErrInvalidZipCode)
}

func TestGivenAValidZipCode_WhenNoAddressesFound_ThenShouldReturnError(t *testing.T) {
	useCase := LoadUseCase()

	var dto GetWeatherByZipCodeDTO
	dto.ZipCode = "00000000"

	_, err := useCase.Execute(dto)

	assert.ErrorIs(t, err, ErrZipCodeNotFound)
}

func TestGivenAValidZipCode_WhenValidAddressFound_ThenShouldReturnWeatherData(t *testing.T) {
	useCase := LoadUseCase()

	var dto GetWeatherByZipCodeDTO
	dto.ZipCode = "99999999"

	output, err := useCase.Execute(dto)

	assert.Nil(t, err)
	assert.IsType(t, float64(0), output.TempC)
	assert.IsType(t, float64(0), output.TempF)
	assert.IsType(t, float64(0), output.TempK)
}
