package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{7},
				{2, 2, 3},
			},
		},
		{
			name: "case2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, combinationSum(tt.args.candidates, tt.args.target))
			// if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			// }
		})
	}
}