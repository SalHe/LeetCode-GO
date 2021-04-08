package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 || len(strs[0]) == 0 {
		return strs[0]
	}

	maxLength := 0
	for maxLength < len(strs[0]) {
		ok := true
		for i := 1; i < len(strs); i++ {
			str := strs[i]
			if maxLength >= len(str) || str[maxLength] != strs[0][maxLength] {
				ok = false
				break
			}
		}
		if ok {
			maxLength++
		} else {
			break
		}
	}
	return strs[0][0:maxLength]
}
