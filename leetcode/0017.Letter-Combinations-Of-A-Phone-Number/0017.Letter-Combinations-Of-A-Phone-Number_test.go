package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question17 struct {
	digits       string
	combinations []string
}

func Test_letterCombinations(t *testing.T) {
	qs := []question17{
		{
			digits:       "",
			combinations: []string{},
		},
		{
			digits:       "23",
			combinations: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, q := range qs {
		assert.ElementsMatch(t, letterCombinations(q.digits), q.combinations)
	}
}
