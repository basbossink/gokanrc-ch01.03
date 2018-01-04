package internal

import (
	"math"
	"testing"
)

const epsilon = 1e-5

var celsiusToFahrenheitTests = []struct {
	celsius            float64
	expectedFahrenheit float64
}{
	{0.0, 32.0},
	{5.0 / 9.0, 33.0},
	{5.0, 41.0},
	{-5.0, 23.0},
	{10.0, 50.0},
	{15.0, 59.0},
	{35.0, 95.0},
}

func TestCelsiusToFahrenheitConversion(t *testing.T) {
	for _, tt := range celsiusToFahrenheitTests {
		actual := ConvertCelsiusToFahrenheit(tt.celsius)
		if epsilon < math.Abs(actual-tt.expectedFahrenheit) {
			t.Fatalf(
				"ConvertCelsiusToFahrenheit(%v). Expected [%v], Actual [%v]",
				tt.celsius,
				tt.expectedFahrenheit,
				actual)
		}

	}
}

var fahrenheitToCelsiusTests = []struct {
	fahrenheit      float64
	expectedCelsius float64
}{
	{32.0, 0.0},
	{33.0, 5.0 / 9.0},
	{41.0, 5.0},
	{23.0, -5.0},
	{50.0, 10.0},
	{59.0, 15.0},
	{95.0, 35.0},
}

func TestFahrenheitToCelsiusConversion(t *testing.T) {
	for _, tt := range fahrenheitToCelsiusTests {
		actual := ConvertFahrenheitToCelsius(tt.fahrenheit)
		if epsilon < math.Abs(actual-tt.expectedCelsius) {
			t.Fatalf(
				"ConvertFahrenheitToCelsius(%v). Expected [%v], Actual [%v]",
				tt.fahrenheit,
				tt.expectedCelsius,
				actual)
		}

	}
}
