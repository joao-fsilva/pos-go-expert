package entity

import (
	"errors"
	"regexp"
)

var ErrInvalidZipCode = errors.New("invalid zipcode")

type ZipCode struct {
	ZipCode string
}

func NewZipCode(cepStr string) (*ZipCode, error) {
	zipCode := &ZipCode{
		ZipCode: cepStr,
	}

	err := zipCode.IsValid()
	if err != nil {
		return nil, err
	}

	return zipCode, nil
}

func (c *ZipCode) IsValid() error {
	regex := regexp.MustCompile(`^\d{8}$`)
	if !regex.MatchString(c.ZipCode) {
		return ErrInvalidZipCode
	}

	return nil
}
