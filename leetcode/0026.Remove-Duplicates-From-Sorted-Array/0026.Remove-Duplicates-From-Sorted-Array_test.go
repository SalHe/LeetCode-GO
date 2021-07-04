package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question26 struct {
	input    []int
	expected []int
}

func Test_removeDuplicates(t *testing.T) {
	qs := []question26{
		{
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, q := range qs {
		var n int

		n = removeDuplicates(q.input)
		assert.Equal(t, q.expected, q.input[:n])

		n = removeDuplicates_version2(q.input)
		assert.Equal(t, q.expected, q.input[:n])
	}
}
