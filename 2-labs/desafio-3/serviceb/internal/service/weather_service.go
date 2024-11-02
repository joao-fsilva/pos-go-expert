package service

import "context"

type WeatherService interface {
	GetWeatherInCelsiusByCity(ctx context.Context, city string) (float64, error)
}
