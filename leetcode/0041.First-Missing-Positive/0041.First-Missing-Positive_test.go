package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bakup := make([]int, len(tt.args.nums))
			copy(bakup, tt.args.nums)
			if got := firstMissingPositive(bakup); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
			copy(bakup, tt.args.nums)
			if got := firstMissingPositive_exchange(bakup); got != tt.want {
				t.Errorf("firstMissingPositive_exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
