package leetcode

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"case1", []int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"case1", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"case2", []int{2, 2, 2, 2, 2, 2, 2}, 3, false},
		{"case2", []int{2, 2, 2, 2, 2, 2, 2}, 2, true},
		{"case3", []int{1}, 0, false},
		{"case3", []int{1}, 1, true},
		{"case4", []int{1, 0, 1, 1, 1}, 0, true},
		{"case4", []int{1, 0, 1, 1, 1}, 1, true},
		{"case4", []int{1, 0, 1, 1, 1}, 2, false},
		{"case5", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2, true},
		{"case6", []int{3, 1}, 1, true},
		{"case6", []int{3, 1}, 3, true},
		{"case6", []int{3, 1}, 0, false},
		{"case7", []int{3, 5, 1}, 1, true},
		{"case8", []int{0, 0, 1, 1, 2, 0}, 2, true},
		{"case9", []int{2, 2, 2, 3, 2, 2, 2}, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
