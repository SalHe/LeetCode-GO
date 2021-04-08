package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question11 struct {
	input  []int
	output int
}

func Test_maxArea(t *testing.T) {

	qs := []question11{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{1, 2, 1},
			output: 2,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.output, maxArea(q.input))
	}

}
