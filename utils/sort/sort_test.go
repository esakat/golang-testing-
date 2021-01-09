package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	// Init
	elements := []int{9, 12, 4, 6, 1, 4, 2, 5, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 12 {
		t.Error("first element should be 12")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}