package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want [][]int
	}{
		{
			name: "case1",
			arg:  3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
