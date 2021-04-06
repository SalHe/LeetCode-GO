package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question6 struct {
	s       string
	numRows int
	output  string
}

func Test_zigZagConversion(t *testing.T) {

	qs := []question6{
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			output:  "PINALSIGYAHRPI",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			output:  "PAHNAPLSIIGYIR",
		},
		{
			s:       "A",
			numRows: 1,
			output:  "A",
		},
		{
			s:       "AB",
			numRows: 1,
			output:  "AB",
		},
	}

	for _, q := range qs {
		assert.Equal(t, q.output, zigZagConversion(q.s, q.numRows))
	}

}
