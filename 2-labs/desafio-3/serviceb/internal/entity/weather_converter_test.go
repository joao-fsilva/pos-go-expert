package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnCelsius_WhenConvertToFahrenheit_ThenShouldConvert(t *testing.T) {
	converter := NewWeatherConverter(25)

	assert.Equal(t, 77.00, converter.ToFahrenheit())
}

func TestGivenAnCelsius_WhenConvertToKelvin_ThenShouldConvert(t *testing.T) {
	converter := NewWeatherConverter(25)

	assert.Equal(t, 298.00, converter.ToKelvin())
}
