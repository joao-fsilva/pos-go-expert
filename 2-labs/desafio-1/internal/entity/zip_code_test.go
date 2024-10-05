package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyZipCode_WhenCreateANewZipCode_ThenShouldReceiveAnError(t *testing.T) {
	cep, err := NewZipCode("")

	assert.Equal(t, "invalid zipcode", err.Error())
	assert.Nil(t, cep)
}

func TestGivenAnInvalidZipCode_WhenCreateANewZipCode_ThenShouldReceiveAnError(t *testing.T) {
	cep, err := NewZipCode("1234567")

	assert.Equal(t, "invalid zipcode", err.Error())
	assert.Nil(t, cep)
}

func TestGivenAnValidZipCode_WhenCreateANewZipCode_ThenShouldReceiveZipCode(t *testing.T) {
	zipCode, err := NewZipCode("12345678")

	assert.Equal(t, "12345678", zipCode.ZipCode)
	assert.Nil(t, err)
}
