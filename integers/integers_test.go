package integers

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if expected != result {
		t.Errorf("Expected: %d. Received: %d", expected, result)
	}
}

func TestSum(t *testing.T) {
	numbers := [3]int{25, 35, 9}
	result := Sum(numbers)
	expected := 69

	if expected != result {
		t.Errorf("Expected: %d. Received: %d", expected, result)
	}
}
