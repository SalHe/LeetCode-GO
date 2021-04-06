package leetcode

func zigZagConversion(s string, numRows int) string {
	length := len(s)
	if length <= 1 || numRows == 1 {
		return s
	}

	cols, rest := length/(numRows-1), length%(numRows-1)
	if rest != 0 {
		cols++
	}

	newStr := make([]byte, length)

	// 虽然是双层循环
	// 但是总循环次数实际是s的长度的线性函数
	p := 0
	for y := 0; y < numRows; y++ {
		for x := 0; x < cols; x++ {
			if x%2 == 0 {
				if x*(numRows-1)+y >= length {
					break
				}
				newStr[p] = s[x*(numRows-1)+y]
				p++
			} else if y != 0 && y != numRows-1 {
				if x*(numRows-1)+numRows-1-y >= length {
					break
				}
				newStr[p] = s[x*(numRows-1)+numRows-1-y]
				p++
			}
		}
	}

	return string(newStr)
}
