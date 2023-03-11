package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		rgs   []rg
		want  int
		want1 int
	}{
		{"case1", []rg{{1, 3, 1}, {2, 6, 1}, {5, 7, 1}, {2, 3, 1}}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solve(tt.rgs)
			if got != tt.want {
				t.Errorf("solve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
