package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	leftCount := (m + n + 1) / 2
	left, right := 0, m

	for left < right {
		i := left + (right-left+1)/2
		j := leftCount - i

		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}

	var nums1LeftMax int
	var nums2LeftMax int
	var nums1RightMin int
	var nums2RightMin int

	i := left
	j := leftCount - i
	if i > 0 {
		nums1LeftMax = nums1[i-1]
	} else {
		nums1LeftMax = math.MinInt64
	}
	if j > 0 {
		nums2LeftMax = nums2[j-1]
	} else {
		nums2LeftMax = math.MinInt64
	}
	if i < m {
		nums1RightMin = nums1[i]
	} else {
		nums1RightMin = math.MaxInt64
	}
	if j < n {
		nums2RightMin = nums2[j]
	} else {
		nums2RightMin = math.MaxInt64
	}

	if (m+n)%2 == 0 {
		return (math.Max(float64(nums1LeftMax), float64(nums2LeftMax)) + math.Min(float64(nums1RightMin), float64(nums2RightMin))) / 2
	} else {
		return math.Max(float64(nums1LeftMax), float64(nums2LeftMax))
	}

}
