package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  int
	}{
		{"case1", []int{7, 6, 5, 1, 3}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.nodes); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
