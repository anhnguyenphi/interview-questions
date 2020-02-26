package the_tortoise_and_the_hare

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// link : https://en.wikipedia.org/wiki/Cycle_detection#Floyd.27s_Tortoise_and_Hare
// link : https://zhu45.org/posts/2017/Jun/18/the-tortoise-and-the-hare/

type node struct {
	Val  int
	Next *node
}

func detectCycle(head *node) *node {
	tortoise := head
	hare := head

	// detect loop
	for hare != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next

		if hare == tortoise {
			// find the start of loop
			for tortoise != head {
				tortoise = tortoise.Next
				head = head.Next
			}
			return tortoise
		}
	}

	return nil
}

func TestCetectCycle_Case1(t *testing.T) {
	n1 := &node{Val: 1}
	n2 := &node{Val: 2}
	n3 := &node{Val: 3}
	n4 := &node{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	result := detectCycle(n1)
	assert.Equal(t, result != nil, true)
	assert.Equal(t, result.Val, 2)
}

func TestCetectCycle_Case2(t *testing.T) {
	n1 := &node{Val: 1}
	n2 := &node{Val: 2}
	n3 := &node{Val: 3}
	n4 := &node{Val: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = nil

	result := detectCycle(n1)
	assert.Equal(t, result == nil, true)
}
