package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question18 struct {
	nums         []int
	target       int
	combinations [][]int
}

func Test_fourSum(t *testing.T) {

	qs := []question18{
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			combinations: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			combinations: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, q := range qs {
		assert.ElementsMatch(t, q.combinations, fourSum(q.nums, q.target))
	}

}
