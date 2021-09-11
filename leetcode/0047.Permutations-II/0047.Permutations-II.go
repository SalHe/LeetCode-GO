package leetcode

import "sort"

// 太久没写了，参考了答案的
func permuteUnique(nums []int) (container [][]int) {
	sort.Ints(nums)
	count := len(nums)
	visit := make([]bool, count)
	cur := make([]int, count)
	icur := 0
	if count == 0 { // 在题目中明确说明不会出现此情况
		return make([][]int, 0)
	}

	var search func(pos int)
	search = func(pos int) {
		if pos == count {
			container = append(container, append([]int(nil), cur...))
			return
		}

		for i, num := range nums {
			if visit[i] || (i > 0 && num == nums[i-1] && visit[i-1]) {
				continue
			}
			cur[icur] = num
			visit[i] = true
			icur++
			search(pos + 1)
			visit[i] = false
			icur--
		}
	}

	search(0)

	return
}
