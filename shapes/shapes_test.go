package shapes

import (
	"testing"
)

func TestPerimiter(t *testing.T) {
	perimiterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, expected: 40.0},
		{name: "Circle", shape: Circle{12}, expected: 75.39822368615503},
		{name: "Triangle", shape: Triangle{10, 10}, expected: 32.3606797749979},
	}

	for _, testCase := range perimiterTests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.shape.Perimiter()

			if result != testCase.expected {
				t.Errorf("%#v. Received: %.2f. Expected: %.2f", testCase.shape, result, testCase.expected)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, testCase := range areaTests {
		result := testCase.shape.Area()

		t.Run(testCase.name, func(t *testing.T) {
			if result != testCase.expected {
				t.Errorf("%#v. Received: %.2f. Expected: %.2f", testCase.shape, result, testCase.expected)
			}
		})
	}
}
