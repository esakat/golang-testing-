package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init
	elements := []int{9, 12, 4, 6, 1, 4, 2, 5, 0}

	// Execution
	fmt.Println(elements)
	BubbleSort(elements)
	fmt.Println(elements)

	// Validation
	if elements[0] != 12 {
		t.Error("first element should be 12")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Init
	elements := []int{10, 8, 5, 3, 2, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 10 {
		t.Error("first element should be 10")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
