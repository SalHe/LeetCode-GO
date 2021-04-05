package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question4 struct {
	Nums1  []int
	Nums2  []int
	Output float64
}

func Test_Mediano(t *testing.T) {
	qs := []question4{
		{
			Nums1:  []int{1, 2, 3},
			Nums2:  []int{3, 4, 5},
			Output: 3.0,
		},
		{
			Nums1:  []int{1, 3},
			Nums2:  []int{2},
			Output: 2,
		},
		{
			Nums1:  []int{1, 2},
			Nums2:  []int{3, 4},
			Output: 2.5,
		},
		{
			Nums1:  []int{0},
			Nums2:  []int{0},
			Output: 0,
		},
		{
			Nums1:  []int{},
			Nums2:  []int{1},
			Output: 1,
		},
		{
			Nums1:  []int{2},
			Nums2:  []int{},
			Output: 2,
		},
	}

	for i := range qs {
		assert.Equal(t, qs[i].Output, findMedianSortedArrays(qs[i].Nums1, qs[i].Nums2))
	}

}
