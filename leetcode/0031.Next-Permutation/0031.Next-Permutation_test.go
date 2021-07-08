package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question31 struct {
	nums []int
	next []int
}

func Test_nextPermutation(t *testing.T) {

	qs := []question31{
		// {
		// 	nums: []int{1, 20, 26, 1, 15, 29, 4, 29, 10, 9, 21, 7, 27, 11, 21, 5, 9, 7, 27, 16, 17, 3, 6, 5, 16, 23, 29, 14, 28, 21, 2, 29, 3, 29, 0, 18, 28, 5, 10, 9, 6, 23, 8, 25, 26, 21, 1, 5, 29, 28, 14, 8, 1, 20, 13, 10},
		// 	next: []int{3, 1, 2},
		// },
		{
			nums: []int{2, 3, 1},
			next: []int{3, 1, 2},
		},
		{
			nums: []int{1, 2, 3},
			next: []int{1, 3, 2},
		},
		{
			nums: []int{3, 2, 1},
			next: []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 5},
			next: []int{1, 5, 1},
		},
		{
			nums: []int{1},
			next: []int{1},
		},
	}

	for _, q := range qs {
		backup := q.nums[:]
		nextPermutation(backup)
		assert.Equal(t, q.next, backup)
	}
}
