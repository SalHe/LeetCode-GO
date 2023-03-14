package leetcode

import (
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	c, area := bigCase1()
	c2, area2 := bigCase2()
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
		{"case3", []int{1}, 1},
		{"case4", []int{0, 9}, 9},
		{"case5", c, area},
		{"case6", c2, area2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
			if got := largestRectangleAreaDp(tt.heights); got != tt.want {
				t.Errorf("largestRectangleAreaDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
