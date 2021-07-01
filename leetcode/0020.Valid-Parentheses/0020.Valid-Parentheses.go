package leetcode

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack, iStack := make([]int32, len(s)/2), -1
	for _, c := range s {
		switch c {
		case '{', '[', '(':
			iStack++
			if iStack >= len(stack) {
				return false
			}
			stack[iStack] = c
		case '}':
			if iStack < 0 || '{' != stack[iStack] {
				return false
			}
			iStack--
		case ']':
			if iStack < 0 || '[' != stack[iStack] {
				return false
			}
			iStack--
		case ')':
			if iStack < 0 || '(' != stack[iStack] {
				return false
			}
			iStack--
		default:
			return false
		}
	}
	return iStack < 0
}

func isValid_Stack(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack, iStack := make([]int32, len(s)/2), -1
	for _, c := range s {
		switch c {
		case '{', '[', '(':
			iStack++
			if iStack >= len(stack) {
				return false
			}
			stack[iStack] = c
			break
		case '}':
			if iStack < 0 || '{' != stack[iStack] {
				return false
			}
			iStack--
			break
		case ']':
			if iStack < 0 || '[' != stack[iStack] {
				return false
			}
			iStack--
			break
		case ')':
			if iStack < 0 || '(' != stack[iStack] {
				return false
			}
			iStack--
			break
		default:
			return false
		}
	}
	return iStack < 0
}
