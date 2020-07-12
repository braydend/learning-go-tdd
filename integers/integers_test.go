package integers

import (
	"reflect"
	"testing"
)

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

func TestSumAllTails(t *testing.T) {
	collectionOne := []int{1, 2, 3, 4, 5}
	collectionTwo := []int{4, 3, 2, 1, 0}

	expected := []int{14, 6}
	result := SumAllTails(collectionOne, collectionTwo)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %d. Received: %d", expected, result)
	}
}
