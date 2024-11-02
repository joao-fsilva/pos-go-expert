package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"weather-zip-code/internal/entity"
	"weather-zip-code/internal/usecase"
)

type WeatherController struct {
	usecase *usecase.GetWeatherByZipCode
}

func NewWeatherController(usecase *usecase.GetWeatherByZipCode) WeatherController {
	return WeatherController{
		usecase: usecase,
	}
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (wc WeatherController) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	zipCode := r.URL.Query().Get("zipcode")

	var dto usecase.GetWeatherByZipCodeDTO
	dto.ZipCode = zipCode

	output, err := wc.usecase.Execute(ctx, dto)

	if err != nil {
		switch {
		case errors.Is(err, entity.ErrInvalidZipCode):
			writeErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		case errors.Is(err, usecase.ErrZipCodeNotFound):
			writeErrorResponse(w, http.StatusNotFound, err.Error())
			return
		default:
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(output)
}