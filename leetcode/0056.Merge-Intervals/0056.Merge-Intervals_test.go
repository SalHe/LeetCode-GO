package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name: "case1",
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "case2",
			intervals: [][]int{
				{1, 4},
				{4, 5},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "case3",
			intervals: [][]int{
				{0, 4},
				{1, 4},
			},
			want: [][]int{
				{0, 4},
			},
		},
		{
			name: "case4",
			intervals: [][]int{
				{2, 3},
				{1, 4},
			},
			want: [][]int{
				{1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
