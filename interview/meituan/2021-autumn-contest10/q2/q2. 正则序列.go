package main

import (
	"fmt"
	"sort"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solve(numbers []int) int {
	sort.Ints(numbers)
	op := 0
	for i, number := range numbers {
		op += abs(number - (i + 1))
	}
	return op
}

// 提交时间：2023-03-10 语言：Go 运行时间：71ms 占用内存：1480 状态：编译正确
func main() {
	var count int
	fmt.Scanf("%d", &count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	fmt.Println(solve(numbers))
}
