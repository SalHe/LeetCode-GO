package leetcode

func groupAnagrams(strs []string) (anagrams [][]string) {
	// 仅包含小写字母

	n := len(strs)
	groups := make([]int, n)
	isAnagram := func(str1, str2 string) bool {
		chars := make([]int, 26)
		for _, c := range str1 {
			chars[c-'a']++
		}
		for _, c := range str2 {
			if chars[c-'a'] <= 0 {
				return false
			}
			chars[c-'a']--
		}
		a := 0
		for _, c := range chars {
			a += c
		}
		return a == 0
	}

	for i, str := range strs {
		groups[i] = i
		for j := 0; j < i; j++ {
			if isAnagram(str, strs[j]) {
				groups[i] = groups[j]
				break
			}
		}

		if groups[i] == i {
			groups[i] = len(anagrams)
			anagrams = append(anagrams, append([]string{}, str))
		} else {
			anagrams[groups[i]] = append(anagrams[groups[i]], str)
		}
	}

	return
}
