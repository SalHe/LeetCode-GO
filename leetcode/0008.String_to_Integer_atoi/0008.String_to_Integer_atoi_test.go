package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question8 struct {
	input  string
	output int
}

func Test_myAtoi(t *testing.T) {
	qs := []question8{
		{
			input:  "42",
			output: 42,
		},
		{
			input:  "-42",
			output: -42,
		},
		{
			input:  "4193 with words",
			output: 4193,
		},
		{
			input:  "words and 987",
			output: 0,
		},
		{
			input:  "-91283472332",
			output: -2147483648,
		},
		{
			input:  "00000-42a1234",
			output: 0,
		},
		{
			input:  "-5-",
			output: -5,
		},
		{
			input:  " + 413",
			output: 0,
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.output, myAtoi(q.input), q.input)
		assert.Equal(t, q.output, myAtoi_Automaton(q.input), q.input)
	}
}
