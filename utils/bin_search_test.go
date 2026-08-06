// Test for the utils package. Covers BinarySearch with table-driven cases
// over ints (boundary conditions, empty/single-element slices, negative
// numbers) and a separate check confirming it works with strings too.
package utils_test

import (
	"testing"

	"example.com/ms-demo-1/utils"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"found in middle", []int{1, 3, 5, 7, 9}, 5, 2},
		{"found at start", []int{1, 3, 5, 7, 9}, 1, 0},
		{"found at end", []int{1, 3, 5, 7, 9}, 9, 4},
		{"not found between elements", []int{1, 3, 5, 7, 9}, 4, -1},
		{"not found below range", []int{1, 3, 5, 7, 9}, -1, -1},
		{"not found above range", []int{1, 3, 5, 7, 9}, 10, -1},
		{"empty slice", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 1, -1},
		{"negative numbers", []int{-9, -5, -1, 0, 3}, -5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.BinarySearch(tt.arr, tt.target); got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d",
					tt.arr, tt.target, got, tt.want)
			}
		})
	}
}

func TestBinarySearchStrings(t *testing.T) {
	arr := []string{"apple", "banana", "cherry", "date"}

	if got := utils.BinarySearch(arr, "cherry"); got != 2 {
		t.Errorf("BinarySearch(%v, %q) = %d, want %d", arr, "cherry", got, 2)
	}
	if got := utils.BinarySearch(arr, "elderberry"); got != -1 {
		t.Errorf("BinarySearch(%v, %q) = %d, want %d", arr, "elderberry", got, -1)
	}
}
