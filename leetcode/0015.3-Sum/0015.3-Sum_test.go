package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question15 struct {
	input  []int
	output [][]int
}

func Test_ThreeSum(t *testing.T) {
	qs := []question15{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			input:  []int{},
			output: [][]int{},
		},
		{
			input:  []int{0},
			output: [][]int{},
		},
		{
			input: []int{1, -1, -1, 0},
			output: [][]int{
				{-1, 0, 1},
			},
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.output, threeSum(q.input))
	}

}
