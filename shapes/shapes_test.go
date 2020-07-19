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
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if expected != result {
			t.Errorf("Received: %.2f. Expected: %.2f", result, expected)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0

		checkArea(t, rectangle, expected)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
