package leetcode

func fullJustify(words []string, maxWidth int) (result []string) {
	length := 0
	start := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		if length+len(word)+(i-start) <= maxWidth {
			length += len(word)
		} else {
			result = append(result, arrange(words[start:i], length, maxWidth, false))

			start = i
			length = 0
			i--
		}
	}

	result = append(result, arrange(words[start:], length, maxWidth, true))

	return
}

func arrange(words []string, length int, maxWidth int, rightAlignment bool) string {
	bytes := make([]byte, maxWidth)
	for i := 0; i < maxWidth; i++ {
		bytes[i] = ' '
	}

	var each, rest int
	if rightAlignment || len(words) == 1 {
		each = 1
		rest = 0
	} else {
		each = (maxWidth - length) / (len(words) - 1)
		rest = (maxWidth - length) % (len(words) - 1)
	}

	p := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		copy(bytes[p:p+len(word)], word)

		p += len(word)
		p += each
		if rest > 0 {
			p++
			rest--
		}
	}

	return string(bytes)
}
