package leetcode

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"case1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"case2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"case3", []int{1, 9, 9}, []int{2, 0, 0}},
		{"case4", []int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
