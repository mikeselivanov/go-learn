// Package utils holds small, generic helper functions used across the
// scratchpad. This file implements BinarySearch, a generic binary search
// over a sorted slice of any ordered type.
package utils

import "cmp"

// BinarySearch returns the index of target in arr, or -1 if not found.
// arr must be sorted in ascending order.
func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return -1
}
