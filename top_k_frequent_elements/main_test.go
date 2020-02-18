package top_k_frequent_elements

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)
// link : https://leetcode.com/problems/top-k-frequent-elements/
// =>  Bucket Sort
// put the same frequency into a bucket and then loop from max to min frequency, append elements of each bucket to arr,
// elements those are from 0 -> k is what we need

func run(arr []int, k int) []int {
	store := map[int]int{}
	freq := map[int][]int{}
	for _, ele := range arr {
		_, ok := store[ele]
		if !ok {
			store[ele] = 1
		} else {
			store[ele] += 1
		}
	}
	for key, value := range store {
		if _, ok := freq[value]; ok {
			freq[value] = append(freq[value], key)
		} else {
			freq[value] = []int{key}
		}
	}

	tmp := make([]int, 0)
	for i := len(arr); i > 0; i-- {
		if val, ok := freq[i]; ok {
			tmp = append(tmp, val...)
		}

		if len(tmp) > k {
			break
		}
	}

	count := math.Min(float64(len(tmp)), float64(k))

	return tmp[:int(count)]
}

func TestCase1(t *testing.T)  {
	assert.Equal(t, run([]int{1,1,1,2,2,3}, 2), []int{1, 2})
}

func TestCase2(t *testing.T)  {
	assert.Equal(t, run([]int{1,1,1,2,2,3,3,3}, 2), []int{1, 3})
}

func TestCase3(t *testing.T)  {
	assert.Equal(t, run([]int{-1,-1}, 1), []int{-1})
}

func TestCase4(t *testing.T)  {
	assert.Equal(t, run([]int{5,3,1,1,1,3,73,1}, 2), []int{1, 3})
}
