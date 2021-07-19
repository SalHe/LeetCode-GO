package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: "case2",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "case3",
			args: args{
				num1: "9133",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
			if got := multiply_stupid(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		a := rand.Uint32()
		b := rand.Uint32()
		c := uint64(a) * uint64(b)
		if int(c) < 0 {
			i--
			continue
		}
		t.Run(fmt.Sprintf("%v * %v = %v", a, b, c), func(t *testing.T) {
			assert.Equal(t, strconv.Itoa(int(c)), multiply(strconv.Itoa(int(a)), strconv.Itoa(int(b))))
		})
	}
}
