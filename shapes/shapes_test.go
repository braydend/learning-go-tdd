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
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, testCase := range areaTests {
		result := testCase.shape.Area()

		if result != testCase.expected {
			t.Errorf("Received: %.2f. Expected: %.2f", result, testCase.expected)
		}
	}
}
