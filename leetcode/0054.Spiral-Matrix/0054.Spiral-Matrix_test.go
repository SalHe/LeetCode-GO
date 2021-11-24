package leetcode

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]int
		wantAns []int
	}{
		{
			name: "3*3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			wantAns: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "3*4",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			wantAns: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spiralOrder(tt.matrix); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spiralOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
