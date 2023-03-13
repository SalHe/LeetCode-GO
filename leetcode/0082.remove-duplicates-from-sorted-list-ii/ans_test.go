package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"case1", []int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{"case2", []int{1, 1, 1, 2, 3}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(BuildList(tt.args)); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
