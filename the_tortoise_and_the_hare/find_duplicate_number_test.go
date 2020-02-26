package the_tortoise_and_the_hare

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// link: https://zhu45.org/posts/2017/Jun/18/the-tortoise-and-the-hare/
// look at table of index and value, we can see a structure of linked list where index i value and value is next pointer
// so if the array has duplicate number mean that the array also has a cycle (idx, val) -> (1, 4), (4, 1)

func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[nums[0]]

	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
	}

	// start of index = 0
	head := 0
	for head != tortoise {
		head = nums[head]
		tortoise = nums[tortoise]
	}

	return head
}

func TestFindDuplicate_Case1(t *testing.T) {
	assert.Equal(t, findDuplicate([]int{1, 3, 4, 2, 2}), 2)
}

func TestFindDuplicate_Case2(t *testing.T) {
	assert.Equal(t, findDuplicate([]int{3, 1, 3, 4, 2}), 3)
}
