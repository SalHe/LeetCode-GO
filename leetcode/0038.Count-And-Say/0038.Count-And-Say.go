package leetcode

// 1 <= n <= 30
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	old := countAndSay(n - 1)
	var fresh []byte
	var last byte = old[0]
	i, count := 1, 1
	for ; i < len(old); i++ {
		if last == 0 || last != old[i] {
			fresh = append(fresh, '0'+byte(count), last)
			count = 1
			last = old[i]
		} else {
			count++
		}
	}
	fresh = append(fresh, '0'+byte(count), last)
	return string(fresh)
}
