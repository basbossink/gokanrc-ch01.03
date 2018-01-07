package internal

import (
	"math"
	"testing"
)

const epsilon = 1e-5

type temperaturePair = struct {
	celsius    float64
	fahrenheit float64
}

var testCases = []temperaturePair{
	{0.0, 32.0},
	{5.0 / 9.0, 33.0},
	{5.0, 41.0},
	{-5.0, 23.0},
	{10.0, 50.0},
	{15.0, 59.0},
	{35.0, 95.0},
}

func runTests(t *testing.T,
	sutName string,
	sut TemperatureConversion,
	inputSelector func(temperaturePair) float64,
	expectedSelector func(temperaturePair) float64) {
	for _, tt := range testCases {
		input := inputSelector(tt)
		expected := expectedSelector(tt)
		actual := sut(input)
		if epsilon < math.Abs(actual-expected) {
			t.Fatalf(
				"%v(%v). Expected [%v], Actual [%v]",
				sutName,
				input,
				expected,
				actual)
		}
	}
}

func TestCelsiusToFahrenheitConversion(t *testing.T) {
	runTests(t,
		"CelsiusToFahrenheit",
		ConvertCelsiusToFahrenheit,
		func(p temperaturePair) float64 { return p.celsius },
		func(p temperaturePair) float64 { return p.fahrenheit })
}

func TestFahrenheitToCelsiusConversion(t *testing.T) {
	runTests(t,
		"FahrenheitToCelsius",
		ConvertFahrenheitToCelsius,
		func(p temperaturePair) float64 { return p.fahrenheit },
		func(p temperaturePair) float64 { return p.celsius })
}
