package internal

import (
	"math"
	"testing"
)

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

const epsilon = 1e-5

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
