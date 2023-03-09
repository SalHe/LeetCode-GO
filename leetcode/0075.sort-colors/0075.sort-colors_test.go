package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"case1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"case2", []int{2, 0, 1}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s := fmt.Sprintf("%v", tt.nums)
			if sortColors(tt.nums); !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
