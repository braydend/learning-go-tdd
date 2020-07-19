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
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, testCase := range areaTests {
		result := testCase.shape.Area()

		if result != testCase.expected {
			t.Errorf("Received: %.2f. Expected: %.2f", result, testCase.expected)
		}
	}
}
