package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question16 struct {
	nums     []int
	target   int
	expected int
}

func Test_threeSumClosest(t *testing.T) {

	qs := []question16{
		{
			nums:     []int{0, 2, 1, -3},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
		{
			nums:     []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			nums:     []int{-1, 2, 1, -4},
			target:   -5,
			expected: -4,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.expected, threeSumClosest(q.nums, q.target), "%v - %v", q.target, q.nums)
	}

}
