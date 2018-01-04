package internal

type TemperatureConversion func(float64) float64

const (
	fahrenheitToCelsiusScalingFactor = 5.0 / 9.0
	zeroCelsiusInFahrenheit          = 32.0
)

func ConvertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius/fahrenheitToCelsiusScalingFactor + zeroCelsiusInFahrenheit
}

func ConvertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - zeroCelsiusInFahrenheit) * fahrenheitToCelsiusScalingFactor
}
