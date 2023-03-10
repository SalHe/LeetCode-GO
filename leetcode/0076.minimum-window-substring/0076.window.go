package leetcode

// 已查看【参考答案】
func minWindow(s string, t string) string {
	needed := make(map[byte]int)
	for _, r := range t {
		needed[byte(r)] += 1
	}

	sLen := len(s)
	filtered := make([]byte, sLen)
	copy(filtered, s)
	index := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		index[i] = i
	}
	for i := sLen - 1; i >= 0; i-- {
		if _, ok := needed[filtered[i]]; !ok {
			filtered = append(filtered[:i], filtered[i+1:]...)
			index = append(index[:i], index[i+1:]...)
		}
	}

	in := make(map[byte]int)
	check := func() bool {
		for ch, count := range needed {
			if count > in[ch] {
				return false
			}
		}
		return true
	}

	ll, rr := -1, -1
	for l, r := 0, 0; l <= r && r < len(filtered); {
		in[filtered[r]]++
		for l <= r && l < sLen && check() {
			if ll == -1 || index[r]-index[l] <= rr-ll {
				ll, rr = index[l], index[r]
			}
			in[filtered[l]]--
			l++
		}
		r++
	}

	if ll == -1 {
		return ""
	}
	return s[ll : rr+1]
}
