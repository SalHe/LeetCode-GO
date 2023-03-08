package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"case1", "0", true},
		{"case2", "e", false},
		{"case3", ".", false},
		{"case4", "12.52", true},
		{"case5", "+12.52", true},
		{"case6", "+.52", true},
		{"case7", "+12.", true},
		{"case8", "+.25", true},
		{"case9", "+.25e333", true},
		{"case10", "+1.25e333", true},
		{"case11", "+1.25e", false},
		{"case12", "005047e+6", true},
		{"case13", "005047e+", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
