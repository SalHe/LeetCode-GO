package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question14 struct {
	strs   []string
	prefix string
}

func Test_longestCommonPrefix(t *testing.T) {

	qs := []question14{
		{
			strs:   []string{"flower", "flow", "flight"},
			prefix: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			prefix: "",
		},
		{
			strs:   []string{"a"},
			prefix: "a",
		},
		{
			strs:   []string{"", ""},
			prefix: "",
		},
		{
			strs:   []string{"ab", "a"},
			prefix: "a",
		},
		{
			strs:   []string{"a", "ac"},
			prefix: "a",
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.prefix, longestCommonPrefix(q.strs))
	}

}
