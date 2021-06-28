package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for p1 := 0; p1 < len(nums)-2; p1++ {
		p2, p3 := p1+1, len(nums)-1
		for p2 < p3 {
			curr := nums[p1] + nums[p2] + nums[p3]
			if math.Abs(float64(curr-target)) < math.Abs(float64(result-target)) {
				result = curr
			}
			if curr < target { // 说明要把result往大的方向靠
				p2++
			} else if curr > target {
				p3--
			} else { // 相等
				return curr
			}
		}
	}
	return result
}

// 这种只能适用于没有负数的情况
//func threeSumClosest(nums []int, target int) int {
//	// 升序排序
//	sort.Ints(nums)
//	for i := range nums {
//		nums[i] *= 3
//	}
//
//	// 二分法定位target应在的位置
//	i, left, right := len(nums)/2, 0, len(nums)-1
//	for left < right && i < len(nums)-1 {
//		if nums[i] > target {
//			right, i = i, left+(right-left)/2
//		} else if nums[i+1] < target {
//			left, i = i, i+(right-left)/2
//		} else {
//			break
//		}
//	}
//
//	// 已经确定位置
//	result := 0
//	p1, p2, need := i, i+1, 3
//	for need > 0 {
//		if p1 < 0 ||
//			(p2 < len(nums) && math.Abs(float64(target-nums[p1])) > math.Abs(float64(target-nums[p2]))) {
//			result += nums[p2]
//			p2++
//		} else {
//			result += nums[p1]
//			p1--
//		}
//
//		need--
//	}
//
//	return result / 3
//}
