package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"case1", "/home/", "/home"},
		{"case2", "/../", "/"},
		{"case3", "/home//foo/", "/home/foo"},
		{"case4", "/a/./b/../../c/", "/c"},
		{"case5", "/", "/"},
		{"case6", "/a/../../b/../c//.//", "/c"},
		{"case7", "/.", "/"},
		{"case8", "/home/foo/./.././bar", "/home/bar"},
		{"case8", "/..", "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
			if got := simplifyPathCharByChar(tt.path); got != tt.want {
				t.Errorf("simplifyPathCharByChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
