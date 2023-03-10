package leetcode

// 已查看【参考答案】
func minWindow(s string, t string) string {
	needed := make(map[byte]int)
	for _, r := range t {
		needed[byte(r)] += 1
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
	for l, r := 0, 0; l <= r && r < len(s); {
		in[s[r]]++
		for l <= r && l < len(s) && check() {
			if ll == -1 || r-l <= rr-ll {
				ll, rr = l, r
			}
			in[s[l]]--
			l++
		}
		r++
	}

	if ll == -1 {
		return ""
	}
	return s[ll : rr+1]
}
