package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		wantAnagrams [][]string
	}{
		{
			name:         "case1",
			args:         []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			wantAnagrams: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:         "empty",
			args:         []string{},
			wantAnagrams: [][]string{},
		},
		{
			name:         "a",
			args:         []string{"a"},
			wantAnagrams: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, anagram := range tt.wantAnagrams {
				sort.Strings(anagram)
			}
			result := groupAnagrams(tt.args)
			for _, strings := range result {
				sort.Strings(strings)
			}
			assert.ElementsMatch(t, tt.wantAnagrams, result)
		})
	}
}
