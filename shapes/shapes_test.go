package shapes

import "testing"

func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimiter(rectangle)
	expected := 40.0

	if expected != result {
		t.Errorf("Received: %.2f. Expected: %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := rectangle.Area()
		expected := 72.0

		if expected != result {
			t.Errorf("Received: %.2f. Expected: %.2f", result, expected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		result := circle.Area()
		expected := 314.1592653589793

		if expected != result {
			t.Errorf("Received: %.2f. Expected: %.2f", result, expected)
		}
	})
}
