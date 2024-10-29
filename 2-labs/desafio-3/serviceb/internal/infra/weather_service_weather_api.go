package infra

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

var ErrRequestWeather = errors.New("an error occurred while processing your request of weather")

type WeatherResponse struct {
	Current CurrentWeather `json:"current"`
}

type CurrentWeather struct {
	TempC float64 `json:"temp_c"`
}

type WeatherServiceWeatherApi struct {
	apiKey string
}

func NewWeatherServiceWeatherApi(apiKey string) *WeatherServiceWeatherApi {
	return &WeatherServiceWeatherApi{
		apiKey: apiKey,
	}
}

func (w *WeatherServiceWeatherApi) GetWeatherInCelsiusByCity(city string) (float64, error) {
	log.Printf("Info: checking the weather of the city of %s", city)

	encodedCity := url.QueryEscape(city)
	urlStr := "http://api.weatherapi.com/v1/current.json?q=" + encodedCity + "&key=" + w.apiKey

	res, err := http.Get(urlStr)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: unexpected status code %d from Wheater API response. URL: %s", res.StatusCode, urlStr)
		return 0.00, ErrRequestWeather
	}

	body, err := io.ReadAll(res.Body)

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)

	return weather.Current.TempC, err
}
