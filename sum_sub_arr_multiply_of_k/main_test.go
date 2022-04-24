package sum_sub_arr_multiply_of_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// link: https://leetcode.com/problems/continuous-subarray-sum/
// (a+(n*x))%x is same as (a%x)
//
// For e.g. in case of the array [23,2,6,4,7] the running sum is [23,25,31,35,42] and the remainders are [5,1,1,5,0]. We got remainder 5 at index 0 and at index 3.
// That means, in between these two indexes we must have added a number which is multiple of the k. Hope this clarifies your doubt :)
func checkSubarraySum(nums []int, k int) bool {
	runningSum := 0
	kv := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		runningSum += nums[i]
		runningSum %= k
		if val, ok := kv[runningSum]; ok && i-val > 1 {
			return true
		}
		kv[runningSum] = i
	}
	return false
}

func TestCheckSubarraySum(t *testing.T) {
	assert.Equal(t, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6), true)
	assert.Equal(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6), true)
	assert.Equal(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 13), false)
}
