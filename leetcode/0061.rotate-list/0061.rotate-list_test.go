package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		k       int
		want    []int
	}{
		{
			name:    "case1",
			numbers: []int{1, 2, 3, 4, 5},
			k:       2,
			want:    []int{4, 5, 1, 2, 3},
		},
		{
			name:    "case2",
			numbers: []int{1, 2},
			k:       0,
			want:    []int{1, 2},
		},
		{
			name:    "case3",
			numbers: []int{1, 2},
			k:       2,
			want:    []int{1, 2},
		},
		{
			name:    "case4",
			numbers: []int{1, 2, 3},
			k:       2000000000,
			want:    []int{2, 3, 1},
		},
		{
			name:    "case5",
			numbers: []int{0, 1, 2},
			k:       4,
			want:    []int{2, 0, 1},
		},
		{
			name:    "case6",
			numbers: []int{1, 2, 3, 4, 5},
			k:       10,
			want:    []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(BuildList(tt.numbers), tt.k).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
