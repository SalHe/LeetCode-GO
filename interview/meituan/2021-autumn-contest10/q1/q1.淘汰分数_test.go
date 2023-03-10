package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		scores []int
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				scores: []int{1, 2, 3, 4, 5, 6},
				x:      2,
				y:      3,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				scores: []int{169, 218, 217, 369, 493, 93, 460, 551, 754, 431, 34, 932, 644, 576, 560, 12, 324, 176, 867, 29, 399, 62, 848, 980, 549, 98, 938, 121, 923, 661, 744, 767, 873, 419, 764, 317, 164, 656, 997, 399, 278, 552, 437, 853, 993, 5, 332, 643, 759, 541, 803, 919, 293, 849, 562, 335, 396, 983, 635, 424, 807, 537, 596, 640, 255, 868, 425, 37, 186, 482, 568, 571, 863, 505, 987, 339, 764, 890, 745, 518, 497, 242, 133, 22, 122, 556, 502, 301, 576, 917, 605, 732, 601, 366, 955, 302, 996, 654, 560, 898},
				x:      45,
				y:      60,
			},
			want: 505,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.scores, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
