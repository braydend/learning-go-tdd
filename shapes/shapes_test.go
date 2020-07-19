package shapes

import "testing"

func TestPerimiter(t *testing.T) {
	result := Perimiter(10.0, 10.0)
	expected := 40.0

	if expected != result {
		t.Errorf("Received: %.2f. Expected: %.2f", expected, result)
	}
}
