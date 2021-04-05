package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Question5 struct {
	Input  string
	Output string
}

func Test_LongestPalindromicSubString(t *testing.T) {

	qs := []Question5{
		{
			Input:  "babad",
			Output: "bab",
		},
		{
			Input:  "cbbd",
			Output: "bb",
		},
		{
			Input:  "a",
			Output: "a",
		},
		{
			Input:  "ac",
			Output: "a",
		},
		{
			Input:  "aacabdkacaa",
			Output: "aca",
		},
	}

	for i := range qs {
		assert.Equal(t, qs[i].Output, longestPalindrome(qs[i].Input))
	}

}
