package main

import "fmt"

// AC

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(solve(s))
}

func solve(str string) int {
	strlen := len(str) // >=1 且 ...
	dp := make([]int, strlen)
	changed := make([]bool, strlen)
	for i := 1; i < strlen; i++ {
		dp[i] = dp[i-1]
		if str[i] == str[i-1] && !changed[i-1] {
			changed[i] = true
			dp[i]++
		}
	}
	return dp[strlen-1]
}
