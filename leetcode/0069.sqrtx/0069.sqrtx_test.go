package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	for i := 0; i < 10000; i++ {
		x := i // rand.Intn(50)
		y := int(math.Sqrt(float64(x)))
		t.Run(fmt.Sprintf("X=%d,Y=%d", x, y), func(t *testing.T) {
			if got := mySqrt(x); got != y {
				t.Errorf("mySqrt() = %v, want %v", got, y)
			}
		})
	}
}
