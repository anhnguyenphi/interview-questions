package k_largest_element_in_array

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func swap(a, b int) (int, int) {
	return b, a
}

func quickSort(nums []int, low, high, k int) int {
	pivot := low
	for i := low; i < high; i++ {
		if nums[i] <= nums[high] {
			nums[i], nums[pivot] = swap(nums[i], nums[pivot])
			pivot++
		}
	}
	nums[pivot], nums[high] = swap(nums[pivot], nums[high])

	count := high - pivot + 1
	if count == k {
		return nums[pivot]
	}

	if count > k {
		return quickSort(nums, pivot+1, high, k)
	}

	return quickSort(nums, low, pivot-1, k-count)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
