package integers

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if expected != result {
		t.Errorf("Expected: %d. Received: %d", expected, result)
	}
}
