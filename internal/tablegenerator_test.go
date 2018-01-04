package internal

import (
	"bytes"
	"strings"
	"testing"
)

func identity(input float64) float64 { return input }

func TestTableGenaration(t *testing.T) {
	var buf bytes.Buffer
	GenerateTable(0.0, 12.0, 3.0, "a\n", "b\n", &buf, identity)
	output := buf.String()
	lines := strings.Split(output, "\n")
	if len(lines) != 6 {
		t.Fatalf(
			"TestTableGeneration: Expected output to have #[%v] lines but got #[%v] lines, output was \n%v",
			5,
			len(lines),
			output)
	}
}
