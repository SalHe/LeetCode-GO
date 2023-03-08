package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{name: "case1", arg: "Hello world", want: 5},
		{name: "case2", arg: "   fly me   to   the moon  ", want: 4},
		{name: "case3", arg: "luffy is still joyboy", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.arg); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
