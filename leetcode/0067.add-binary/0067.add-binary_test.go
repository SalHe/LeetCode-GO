package leetcode

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "case1", a: "11", b: "1", want: "100"},
		{name: "case2", a: "1010", b: "1011", want: "10101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
