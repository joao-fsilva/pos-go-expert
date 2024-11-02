package service

import (
	"context"
	"weather-zip-code/internal/entity"
)

type ZipCodeService interface {
	GetCityByZipCode(ctx context.Context, zipCode *entity.ZipCode) (string, error)
}
