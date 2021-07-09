package controls

import (
	"testing"
)

func TestParser(t *testing.T) {
	separator := "\r"

	data, err := Parser("./links.txt", separator)

	result := len(data)
	expected := 2

	if expected != result || err != nil {
		t.Errorf("Should be %d but got %d ", expected, result)
	}
}
