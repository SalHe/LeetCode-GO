package leetcode

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	ans := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := l2 - 1; j >= 0; j-- {
			v := n1*(num2[j]-'0') + ans[i+j+1]
			ans[i+j+1] = v % 10
			ans[i+j] += v / 10
		}
	}

	i := 0
	for ; i < l1+l2; i++ {
		if ans[i] != 0 {
			break
		}
	}
	if i == l1+l2 {
		return "0"
	}
	s := i
	for ; i < l1+l2; i++ {
		ans[i] += '0'
	}
	return string(ans[s:])
}

func multiply_stupid(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 < l2 {
		l1, l2 = l2, l1
		num1, num2 = num2, num1
	}

	if l2 == 1 {
		if num2 == "0" {
			return "0"
		} else if num2 == "1" {
			return num1
		}

		ans := make([]byte, l1+1) // 一个数乘以另外一个十以内的树最大扩大一个数位
		var in uint8
		n2 := num2[0] - '0'
		for i := 0; i < l1; i++ {
			num := (num1[l1-1-i]-'0')*n2 + in
			in = num / 10
			ans[l1-i] = num%10 + '0'
		}

		if in != 0 {
			ans[0] = in + '0'
			return string(ans)
		} else {
			return string(ans[1:])
		}
	} else {
		ans := "0"
		zeros := ""
		for i := 0; i < l2; i++ {
			r1 := multiply(num1, string(num2[l2-1-i])) + zeros
			zeros += "0"
			i++

			r2 := "0"
			if i < l2 {
				r2 = multiply(num1, string(num2[l2-1-i])) + zeros
				zeros += "0"
			}

			ans = add(ans, r1)
			ans = add(ans, r2)
		}
		return ans
	}
}

func add(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	longer := l2
	if l1 > l2 {
		longer = l1
	}
	var in uint8 = 0
	result := make([]byte, longer+1)
	for i := 1; i <= longer; i++ {
		var a, b uint8 // = 0
		if i <= l1 {
			a = num1[l1-i] - '0'
		}
		if i <= l2 {
			b = num2[l2-i] - '0'
		}
		num := a + b + in
		result[longer+1-i] = num%10 + '0'
		in = num / 10
	}
	if in == 1 {
		result[0] = '1'
		return string(result)
	} else {
		return string(result[1:])
	}
}
