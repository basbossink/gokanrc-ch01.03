package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/basbossink/gokarc-ch01-1.03/internal"
)

const (
	showCelsiusHelp    = "Show the Celsius to Fahrenheit table"
	showFahrenheitHelp = "Show the Fahrenheit to Celsius table"
	shorthand          = " (shorthand)"
	lower              = 0.0
	upper              = 300.0
	step               = 20.0
	headerLine         = `
| Temperature conversion table |
+------------------------------+
`
	celsiusFooterLine      = "+-------------+----------------|\n"
	fahrenheitFooterLine   = "+----------------+-------------|\n"
	celsiusTableHeading    = "|  Celsius    |  Fahrenheit    |\n"
	fahrenheitTableHeading = "|  Fahrenheit    |  Celsius    |\n"
	celsiusRowFormat       = "| %11.0f | %14.1f |\n"
	fahrenheitRowFormat    = "| %14.0f | %11.1f |\n"
	celsiusTableHeader     = headerLine + celsiusTableHeading + celsiusFooterLine
	fahrenheitTableHeader  = headerLine + fahrenheitTableHeading + fahrenheitFooterLine
)

var (
	showCelsiusTable    = flag.Bool("c", false, showCelsiusHelp+shorthand+".")
	showFahrenheitTable = flag.Bool("f", false, showFahrenheitHelp+shorthand+".")
)

func init() {
	flag.BoolVar(showCelsiusTable, "celsius", false, showCelsiusHelp+".")
	flag.BoolVar(showFahrenheitTable, "fahrenheit", false, showFahrenheitHelp+".")
}

func main() {
	flag.Parse()
	both := *showCelsiusTable && *showFahrenheitTable
	if *showCelsiusTable {
		internal.GenerateTable(
			lower,
			upper,
			step,
			celsiusTableHeader,
			celsiusRowFormat,
			os.Stdout,
			internal.ConvertCelsiusToFahrenheit)
	}
	if both {
		fmt.Println("")
	}
	if *showFahrenheitTable {
		internal.GenerateTable(
			lower,
			upper,
			step,
			fahrenheitTableHeader,
			fahrenheitRowFormat,
			os.Stdout,
			internal.ConvertFahrenheitToCelsius)
	}
}
