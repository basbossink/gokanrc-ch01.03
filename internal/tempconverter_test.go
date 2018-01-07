package internal

import (
	"math"
	"testing"
	"testing/quick"
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

func approximatelyEqual(actual, expected float64) bool {
	return math.Abs(actual-expected) < epsilon
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
		if !approximatelyEqual(actual, expected) {
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

func TestConversionBackAndForth(t *testing.T) {
	f := func(randomInput float64) bool {
		converted := ConvertCelsiusToFahrenheit(ConvertFahrenheitToCelsius(randomInput))
		return approximatelyEqual(converted, randomInput)
	}
	if err := quick.Check(f, &quick.Config{MaxCountScale: 1e-3}); err != nil {
		t.Error(err)
	}
}
