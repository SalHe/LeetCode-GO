package main

import (
	"fmt"
)

// 提交时间：2023-03-10 语言：Go 运行时间：2001ms 占用内存：908 状态：编译错误
func main() {
	var nodesCount int
	fmt.Scanf("%d\n", &nodesCount)
	lnrNodes := make([]int, nodesCount)
	for i := 0; i < nodesCount; i++ {
		fmt.Scanf("%d", &lnrNodes[i])
	}
	fmt.Println(solve(lnrNodes))
}

func solve(nodes []int) int {
	return solveRecursive(nodes, 0)
}

func solveRecursive(nodes []int, parent int) int {
	if len(nodes) == 1 {
		return nodes[0] * parent
	}
	sum := min(
		nodes[0]*parent+solveRecursive(nodes[1:], nodes[0]),
		nodes[len(nodes)-1]*parent+solveRecursive(nodes[:len(nodes)-1], nodes[len(nodes)-1]),
	)
	for i := 1; i < len(nodes)-1; i++ {
		sum = min(sum,
			nodes[i]*parent+solveRecursive(nodes[:i], nodes[i])+solveRecursive(nodes[i+1:], nodes[i]))
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
