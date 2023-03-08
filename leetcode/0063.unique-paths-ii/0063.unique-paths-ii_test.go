package leetcode

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "case1",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			name: "case2",
			grid: [][]int{
				{0, 1},
				{0, 0},
			},
			want: 1,
		},
		{
			name: "case3",
			grid: [][]int{
				{1},
			},
			want: 0,
		},
		{
			name: "case4",
			grid: [][]int{
				{0, 0},
			},
			want: 1,
		},
		{
			name: "case5",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			want: 0,
		},
		{
			name: "case6",
			grid: [][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.grid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
