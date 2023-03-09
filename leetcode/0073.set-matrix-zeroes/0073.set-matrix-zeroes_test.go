package leetcode

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			"case1",
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			"case2",
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			"case3",
			[][]int{
				{1, 0},
			},
			[][]int{
				{0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s := fmt.Sprintf("%v", tt.matrix)
			if setZeroes(tt.matrix); !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("simplifyPathCharByChar() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
