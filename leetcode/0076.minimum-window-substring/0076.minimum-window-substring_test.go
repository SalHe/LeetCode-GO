package leetcode

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "case2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "case3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "case4",
			args: args{
				s: "abc",
				t: "cba",
			},
			want: "abc",
		},
		{
			name: "case5",
			args: args{
				s: "aab",
				t: "aab",
			},
			want: "aab",
		},
		{
			name: "case6",
			args: args{
				s: "cabeca",
				t: "cae",
			},
			want: "eca",
		},
		{
			name: "case7",
			args: args{
				s: "cabwefgewcwaefgcf",
				t: "cae",
			},
			want: "cwae",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
