package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			"case1",
			[]int{1, 2, 3},
			[][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO 不按顺序检查
			gotResult := subsets(tt.nums)
			assert.ElementsMatch(t, gotResult, tt.want, "subsets() = %v, want %v", gotResult, tt.want)
		})
	}
}
