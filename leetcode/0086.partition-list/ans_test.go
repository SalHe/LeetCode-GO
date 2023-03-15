package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		name string
		args []int
		x    int
		want []int
	}{
		{
			name: "case1",
			args: []int{1, 4, 3, 2, 5, 2},
			x:    3,
			want: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name: "case2",
			args: []int{2, 1},
			x:    2,
			want: []int{1, 2},
		},
		{
			name: "case3",
			args: []int{1, 4, 3, 2, 5, 2},
			x:    3,
			want: []int{1, 2, 2, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(BuildList(tt.args), tt.x).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
			if got := partitionNoNew(BuildList(tt.args), tt.x).ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionNoNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
