package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question30 struct {
	s      string
	words  []string
	starts []int
}

func Test_findSubstring(t *testing.T) {
	qs := []question30{
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "good"},
			starts: []int{8},
		},
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			starts: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			starts: []int{},
		},
		{
			s:      "wordgoodgoodgoodbestwordxxxx",
			words:  []string{"word", "good", "best", "word"},
			starts: []int{},
		},
		{
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			starts: []int{6, 9, 12},
		},
	}
	for _, q := range qs {
		assert.ElementsMatch(t, q.starts, findSubstring(q.s, q.words), `"%v" contains %v`, q.s, q.words)
	}
}
