package services

import "testing"

func TestSort(t *testing.T) {
	// Init
	elements := []int{10, 8, 5, 3, 2, 0}

	// Execution
	Sort(elements)

	// Validation
	if elements[0] != 10 {
		t.Error("first element should be 10")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}