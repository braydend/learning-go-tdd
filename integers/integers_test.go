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

	t.Run("Sums array of 3 numbers", func(t *testing.T) {
		numbers := []int{25, 35, 9}
		result := Sum(numbers)
		expected := 69

		if expected != result {
			t.Errorf("Expected: %d. Received: %d", expected, result)
		}
	})

	t.Run("Sums array of any length of numbers", func(t *testing.T) {
		numbers := []int{6, 10, 33, 2, 9, 9}
		result := Sum(numbers)
		expected := 69

		if expected != result {
			t.Errorf("Expected: %d. Received: %d", expected, result)
		}
	})
}
