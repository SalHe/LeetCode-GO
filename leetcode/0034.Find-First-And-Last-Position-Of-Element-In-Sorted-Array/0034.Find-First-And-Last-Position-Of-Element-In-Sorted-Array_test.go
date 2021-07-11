package leetcode

import (
	"reflect"
	"testing"
)

// 尝试一下GoLand自动生成的测试代码好不好用
func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{2, 2},
				target: 2,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
		{
			name: "group1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "group2",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 3, 4, 5, 9},
				target: 3,
			},
			want: []int{2, 5},
		},
		{
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "single",
			args: args{
				nums:   []int{0, 0, 1, 2, 3, 3, 4},
				target: 2,
			},
			want: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
			if got := searchRange_linear(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
