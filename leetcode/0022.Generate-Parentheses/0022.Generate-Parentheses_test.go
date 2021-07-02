package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question22 struct {
	n      int
	result []string
}

func Test_generateParenthesis(t *testing.T) {
	qs := []question22{
		{
			n:      1,
			result: []string{"()"},
		},
		{
			n:      3,
			result: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:      4,
			result: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}

	for _, q := range qs {
		assert.ElementsMatch(t, q.result, generateParenthesis(q.n), "n = %v", q.n)
		assert.ElementsMatch(t, q.result, generateParenthesis_version1(q.n), "n = %v", q.n)
	}
}
