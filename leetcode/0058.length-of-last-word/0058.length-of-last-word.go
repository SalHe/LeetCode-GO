package leetcode

func lengthOfLastWord(s string) int {
	end := -1
	start := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch != ' ' {
			if end == -1 {
				end = i
			}
		} else if end != -1 {
			start = i + 1
			break
		}
	}

	return end - start + 1
}
