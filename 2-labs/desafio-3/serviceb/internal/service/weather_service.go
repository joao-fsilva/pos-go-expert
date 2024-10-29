package service

type WeatherService interface {
	GetWeatherInCelsiusByCity(city string) (float64, error)
}
