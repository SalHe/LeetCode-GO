package leetcode

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m > n {
		m, n = n, m
	}
	// m<n

	up, down := 1.0, 1.0
	for i := n; i <= m+n-2; i++ {
		up *= float64(i)
	}
	for j := 1; j <= m-1; j++ {
		down *= float64(j)
	}
	return int(up / down)

	// up := uint64(1)
	// for i := n; i <= m+n-2; i++ {
	// 	up *= uint64(i)
	// 	println(up)
	// }
	// for j := 1; j <= m-1; j++ {
	// 	up /= uint64(j)
	// }
	// return int(up)
}

// ---------------------------------------------------

func uniquePathsByConquer(m int, n int) int {
	// 1 <= m, n <= 100
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	return solve(m, n, cache)
}

func solve(m int, n int, cache [][]int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	if n == 2 { // 针对性优化
		return m
	}
	if m == 2 {
		return n
	}

	if cache[m-1][n-1] != 0 {
		return cache[m-1][n-1]
	}
	a := cache[m-2][n-1]
	if a == 0 {
		a = solve(m-1, n, cache)
	}
	b := cache[m-1][n-2]
	if b == 0 {
		b = solve(m, n-1, cache)
	}
	cache[m-1][n-1] = a + b
	return cache[m-1][n-1]
}
