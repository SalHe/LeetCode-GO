package leetcode

import "math"

func getPermutation(n int, k int) string {
	// 已知 1<=n<=9
	factors := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	digits := make([]byte, n)
	for i := n - 2; i >= 0; i-- {
		factors[i] = factors[i+1] * (n - i - 1)
	}
	for i := 0; i < len(digits); i++ {
		digits[i] = byte('0' + i + 1)
	}

	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		index := int(math.Ceil(float64(k)/float64(factors[i]))) - 1
		bytes[i] = digits[index]
		digits = append(digits[:index], digits[index+1:]...)
		k -= index * factors[i]
	}
	return string(bytes)
}

func getPermutationRecursion(n int, k int) string {
	visited := make([]bool, n+1) // 余一个空间出来，等会就不用 n-1 进行索引了
	bytes := make([]byte, n)
	index := 0
	return solve(n, visited, k, 0, bytes, &index)
}

func solve(n int, visited []bool, k int, current int, bytes []byte, index *int) string {
	if current == n {
		*index++
		if *index == k {
			return string(bytes)
		} else {
			return ""
		}
	}
	for i := 1; i <= n; i++ {
		if !(visited[i]) {
			visited[i] = true
			bytes[current] = byte('0' + i)
			res := solve(n, visited, k, current+1, bytes, index)
			if len(res) == n {
				return res
			}
			visited[i] = false
		}
	}
	return ""
}
