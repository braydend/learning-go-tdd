package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if expected != result {
		t.Errorf("Expected: %q. Received: %q", expected, result)
	}
}
