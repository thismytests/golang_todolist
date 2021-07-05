package openLinksInBrowser

import (
	"testing"
)

func TestParser(t *testing.T) {
	data := parser("./links.txt")
	result := len(data)
	expected := 2

	if expected != result {
		t.Errorf("Should be %d but got %d ", expected, result)
	}
}
