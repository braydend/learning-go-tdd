package shapes

import "testing"

func TestPerimiter(t *testing.T) {
	result := Perimiter(10.0, 10.0)
	expected := 40.0

	if expected != result {
		t.Errorf("Received: %.2f. Expected: %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	result := Area(12.0, 6.0)
	expected := 72.0

	if expected != result {
		t.Errorf("Received: %.2f. Expected: %.2f", result, expected)
	}
}
