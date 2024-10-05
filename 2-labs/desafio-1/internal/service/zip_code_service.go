package service

import (
	"weather-zip-code/internal/entity"
)

type ZipCodeService interface {
	GetCityByZipCode(zipCode *entity.ZipCode) (string, error)
}
