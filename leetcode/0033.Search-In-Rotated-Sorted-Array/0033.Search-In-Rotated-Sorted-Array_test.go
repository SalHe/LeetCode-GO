package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question33 struct {
	nums   []int
	target int
	pos    int
}

func Test_search(t *testing.T) {
	qs := []question33{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			pos:    4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			pos:    -1,
		},
		{
			nums:   []int{1},
			target: 0,
			pos:    -1,
		},
		{
			nums:   []int{1},
			target: 1,
			pos:    0,
		},
		{
			nums:   []int{1, 3},
			target: 1,
			pos:    0,
		},
		{
			nums:   []int{1, 3},
			target: 3,
			pos:    1,
		},
		{
			nums:   []int{3, 1},
			target: 1,
			pos:    1,
		},
	}
	for _, q := range qs {
		assert.Equal(t, q.pos, search(q.nums, q.target))
	}
}
