package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question20 struct {
	str     string
	isValid bool
}

func Test_isValid(t *testing.T) {
	qs := []question20{
		{str: "()", isValid: true},
		{str: "()[]{}", isValid: true},
		{str: "(]", isValid: false},
		{str: "([)]", isValid: false},
		{str: "{[]}", isValid: true},
	}
	for _, q := range qs {
		assert.Equal(t, q.isValid, isValid(q.str), q.str)
		assert.Equal(t, q.isValid, isValid_Stack(q.str), q.str)
	}
}
