package usecase

import (
	"context"
	"errors"
	"weather-zip-code/internal/entity"
	"weather-zip-code/internal/service"
)

var ErrZipCodeNotFound = errors.New("can not find zipcode")

type GetWeatherByZipCodeDTO struct {
	ZipCode string `json:"zip_code"`
}

type GetWeatherByZipCodeOutputDTO struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type GetWeatherByZipCode struct {
	zipCodeService service.ZipCodeService
	weatherService service.WeatherService
}

func NewGetWeatherByZipCode(zipCodeService service.ZipCodeService, weatherService service.WeatherService) *GetWeatherByZipCode {
	return &GetWeatherByZipCode{
		zipCodeService: zipCodeService,
		weatherService: weatherService,
	}
}

func (g *GetWeatherByZipCode) Execute(ctx context.Context, input GetWeatherByZipCodeDTO) (GetWeatherByZipCodeOutputDTO, error) {
	output := GetWeatherByZipCodeOutputDTO{}

	zipCode, err := entity.NewZipCode(input.ZipCode)

	if err != nil {
		return output, err
	}

	city, err := g.zipCodeService.GetCityByZipCode(ctx, zipCode)
	if err != nil {
		return output, err
	}

	if city == "" {
		return output, ErrZipCodeNotFound
	}

	weather, err := g.weatherService.GetWeatherInCelsiusByCity(city)
	if err != nil {
		return output, err
	}

	converter := entity.WeatherConverter{
		Celsius: weather,
	}

	output.City = city
	output.TempC = weather
	output.TempF = converter.ToFahrenheit()
	output.TempK = converter.ToKelvin()

	return output, err
}
