package leetcode

func addBinary(a string, b string) string {
	// 1 <= a.length, b.length <= 10^4
	la, lb := len(a), len(b)
	if la < lb {
		a, b = b, a
		la, lb = lb, la
	}

	bytes := make([]byte, la+1)
	in := byte(0)
	for i := 0; i < la; i++ {
		digitA := a[la-1-i] - '0'
		digitB := byte(0)
		if i < lb {
			digitB = b[lb-i-1] - '0'
		}

		sum := digitA + digitB + in
		digit := sum % 2
		in = sum / 2
		bytes[la-i] = '0' + digit
	}
	bytes[0] = '0' + in

	return string(bytes[1-in:])
}
