package incremental_array

import (
	"fmt"
	"testing"
)

// Find the biggest negative and smallest positive num in sorted array
// time complexity: O(logN)
func BinarySearch(arr []int, left, right int) (int, int) {
	if left > right {
		return -1, -1
	}
	mid := (left + right) / 2
	if arr[mid] == 0 {
		neg, _ := BinarySearch(arr, left, mid-1)
		_, pos := BinarySearch(arr, mid+1, right)
		return neg, pos
	}

	if arr[mid] > 0 {
		if mid == 0 {
			return -1, mid
		}
		if arr[mid-1] <= 0 {
			neg, _ := BinarySearch(arr, left, mid-1)
			return neg, mid
		}
		return BinarySearch(arr, left, mid-1)
	}

	// arr[mid] is negative
	if mid == len(arr)-1 {
		return mid, -1
	}
	if arr[mid+1] >= 0 {
		_, pos := BinarySearch(arr, mid+1, right)
		return mid, pos
	}
	return BinarySearch(arr, mid+1, right)
}

// return neg, pos index
func FindIndices(arr []int) (int, int) {
	return BinarySearch(arr, 0, len(arr)-1)
}

func TestFindIndices(t *testing.T) {
	fmt.Println(FindIndices([]int{-10, -5, -1, 1, 4, 6, 10, 100}))
	fmt.Println(FindIndices([]int{-10, -5, -1, 1, 0, 0, 4, 6, 10, 100}))
	fmt.Println(FindIndices([]int{0, 0, 4, 6, 10, 100}))
	fmt.Println(FindIndices([]int{-10, -5, -1, 0, 0}))
}
