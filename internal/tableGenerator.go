package internal

import (
	"fmt"
	"io"
)

func GenerateTable(
	lower float64,
	upper float64,
	step float64,
	header string,
	rowFormat string,
	writer io.Writer,
	conversion TemperatureConversion) {
	fromTemp := lower
	fmt.Fprintf(writer, header)
	for fromTemp < upper {
		toTemp := conversion(fromTemp)
		fmt.Fprintf(writer, rowFormat, fromTemp, toTemp)
		fromTemp += step
	}
}
