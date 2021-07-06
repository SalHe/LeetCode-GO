package leetcode

func findSubstring(s string, words []string) []int {
	// NOTICE: words中的字符串长度相同

	var result []int

	wordsCount := len(words)
	wordsMap := make(map[string]int, wordsCount)
	for i, word := range words {
		// 实际value目前没有用上，不过可能对wordsCount有帮助，暂时先不考虑了
		wordsMap[word] = i
	}

	if wordsCount > 0 {
		wordLen := len(words[0])
		i, upper := 0, len(s)-wordLen*wordsCount
		for i <= upper {

			allExist := true

			// 分割窗口内字符串的单词
			// 因为目标单词的长度是固定的
			// 所以分法也是固定的
			hashCount := make(map[string]int, wordsCount)
			finalStart := len(s)
			for iWord := wordsCount - 1; iWord >= 0; iWord-- {
				start := i + iWord*wordLen
				if start < finalStart {
					finalStart = start
				}
				wordSplit := s[start : start+wordLen]
				if _, ok := wordsMap[wordSplit]; !ok {
					break

					// 不存在此单词，那么窗口直接可以滑到该单词之后了（这样会漏解，可能需要另外处理才行）
					// i = start + 1
					// goto nextPos
				}

				if count, ok := hashCount[wordSplit]; ok {
					hashCount[wordSplit] = count + 1
				} else {
					hashCount[wordSplit] = 1
				}
			}

			// 现在只要判断是否所有单词都存在即可
			// 并找到起始位置

			for _, word := range words {
				count, ok := hashCount[word]
				if !ok || count <= 0 {
					allExist = false
					break
				}
				hashCount[word] = count - 1
			}

			if allExist {
				result = append(result, finalStart)
			}

			i++

			// nextPos:
		}
	}

	return result
}

// findSubstring_version2 中在分割单词并计数的时候，遇到不存在的单词其实可以直接跳过，不过本代码中没处理
func findSubstring_version2(s string, words []string) []int {
	// NOTICE: words中的字符串长度相同

	var result []int

	wordsCount := len(words)
	if wordsCount > 0 {
		wordLen := len(words[0])
		for i, n := 0, len(s)-wordLen*wordsCount; i <= n; i++ {
			// 分割窗口内字符串的单词
			// 因为目标单词的长度是固定的
			// 所以分法也是固定的
			hashCount := make(map[string]int, wordsCount)
			finalStart := len(s)
			for iWord := wordsCount - 1; iWord >= 0; iWord-- {
				start := i + iWord*wordLen
				if start < finalStart {
					finalStart = start
				}
				wordSplit := s[start : start+wordLen]

				if count, ok := hashCount[wordSplit]; ok {
					hashCount[wordSplit] = count + 1
				} else {
					hashCount[wordSplit] = 1
				}
			}

			// 现在只要判断是否所有单词都存在即可
			// 并找到起始位置
			allExist := true

			for _, word := range words {
				count, ok := hashCount[word]
				if !ok || count <= 0 {
					allExist = false
					break
				}
				hashCount[word] = count - 1
			}

			if allExist {
				result = append(result, finalStart)
			}
		}
	}

	return result
}

// findSubstring_version1 里面多维护了一个hashPos，其实没有必要用它。只会徒增比较次数和存储代价。
func findSubstring_version1(s string, words []string) []int {
	// NOTICE: words中的字符串长度相同

	var result []int

	wordsCount := len(words)
	hashPos := make(map[string]int, wordsCount) // 这个map可以复用（不需重新make），因为hashCount不满足是不会认为成功的
	if wordsCount > 0 {
		wordLen := len(words[0])
		for i, n := 0, len(s)-wordLen*wordsCount; i <= n; i++ {
			// 分割窗口内字符串的单词
			// 因为目标单词的长度是固定的
			// 所以分法也是固定的
			// hashPos := make(map[string]int) // 这个map可以复用（不需重新make），因为hashCount不满足是不会认为成功的
			hashCount := make(map[string]int, wordsCount)
			for iWord := wordsCount - 1; iWord >= 0; iWord-- {
				start := i + iWord*wordLen
				wordSplit := s[start : start+wordLen]
				hashPos[wordSplit] = start

				if count, ok := hashCount[wordSplit]; ok {
					hashCount[wordSplit] = count + 1
				} else {
					hashCount[wordSplit] = 1
				}
			}

			// 现在只要判断是否所有单词都存在即可
			// 并找到起始位置
			allExist, start := true, len(s)
			for _, word := range words {
				p, ok := hashPos[word]
				if !ok {
					allExist = false
					break
				} else {
					count := hashCount[word]
					if count <= 0 {
						allExist = false
						break
					}
					hashCount[word] = count - 1
				}

				if p < start {
					start = p
				}
			}

			if allExist {
				result = append(result, start)
			}
		}
	}

	return result
}
