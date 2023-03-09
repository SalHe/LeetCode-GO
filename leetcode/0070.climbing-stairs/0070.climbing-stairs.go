package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 2
	c := a + b

	for i := 3; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
