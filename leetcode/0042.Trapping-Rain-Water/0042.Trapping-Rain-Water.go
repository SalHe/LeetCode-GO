package leetcode

func trap(height []int) int {
	v := 0
	for i := 0; i < len(height)-1; i++ {
		if height[i] == 0 {
			continue
		}
		j := i + 1
		h := 0
		e := len(height) - 1
		for ; j < len(height); j++ {
			if height[j] >= height[i] {
				break
			} else if height[j] >= h {
				h = height[j]
				e = j
			}
		}

		if j < len(height) {
			h = height[i]
			e = j
		}
		for k := i + 1; k < e; k++ {
			v += h - height[k]
		}
		i = e - 1 // 其实是希望下一次循环时i=e
	}
	return v
}
