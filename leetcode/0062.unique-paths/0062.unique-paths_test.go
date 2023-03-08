package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{"case1", 2, 4, 4},
		{"case2", 3, 7, 28},
		{"case3", 3, 3, 6},
		{"case4", 3, 2, 3},
		{"case5", 51, 8, 264385836},
		{"case6", 13, 23, 548354040},
		{"case7", 16, 16, 155117520},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsByConquer(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePathsByConquer() = %v, want %v", got, tt.want)
			}
		})
	}
}
