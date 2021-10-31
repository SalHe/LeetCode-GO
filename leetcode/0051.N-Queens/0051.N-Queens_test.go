package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name         string
		size         int
		wantSolution [][]string
	}{
		{
			name: "4*4",
			size: 4,
			wantSolution: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name:         "1*1",
			size:         1,
			wantSolution: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantSolution, solveNQueens(tt.size))
		})
	}
}
