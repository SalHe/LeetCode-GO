package leetcode

func permute(nums []int) (container [][]int) {
	// 排列的长度是确定的，排列的情况也是可以计算出来的
	count := len(nums)
	if count == 0 { // 在题目中明确说明不会出现此情况
		return make([][]int, 0)
	}

	all := 1
	for i := 1; i <= count; i++ {
		all *= i
	}

	container = make([][]int, all)
	for i := 0; i < all; i++ {
		container[i] = make([]int, count)
	}

	icurr := 0
	var search func(pos int)
	search = func(pos int) {
		if pos == count {
			copy(container[icurr], nums)
			icurr++
			return
		}

		for i := pos; i < count; i++ {
			nums[i], nums[pos] = nums[pos], nums[i]
			search(pos + 1)
			nums[i], nums[pos] = nums[pos], nums[i]
		}
	}

	search(0)

	return
}
