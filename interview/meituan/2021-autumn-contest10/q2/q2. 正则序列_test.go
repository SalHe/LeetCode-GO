package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{"case1", []int{3, 0, 3, 1, 3}, 5},
		{"case2", []int{-1, 2, 3, 10, 100}, 103},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.numbers); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
