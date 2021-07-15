package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	qs := []string{
		"1",
		"11",
		"21",
		"1211",
		"111221",
		"312211",
		"13112221",
		"1113213211",
		"31131211131221",
		"13211311123113112211",
	}
	for i, q := range qs {
		assert.Equal(t, q, countAndSay(i+1))
	}
}
