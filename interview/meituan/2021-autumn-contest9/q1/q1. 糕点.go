package main

import (
	"fmt"
	"strings"
)

func main() {
	out := &strings.Builder{}
	for {
		var maxCakesCount, nowCakes, weight1, weight2 int
		_, err := fmt.Scanf("%d", &maxCakesCount)
		if err != nil {
			break
		}
		fmt.Scanf("%d%d%d", &nowCakes, &weight1, &weight2)
		fmt.Scanf("\n")
		weights := make([]int, nowCakes)
		for i := 0; i < nowCakes; i++ {
			fmt.Scanf("%d", &weights[i])
		}
		fmt.Scanf("\n")

		if solve(maxCakesCount, weights, weight1, weight2) {
			out.WriteString("YES")
		} else {
			out.WriteString("NO")
		}
		out.WriteString("\n")
	}
	println(out.String())
}

func solve(maxCakesCount int, weights []int, weight1 int, weight2 int) bool {
	min, max := 99999999999, -1
	minCount, maxCount := 0, 0
	for _, weight := range weights {
		if weight > max {
			max = weight
			maxCount = 0
		}
		if weight < min {
			min = weight
			minCount = 0
		}

		if weight == min {
			minCount++
		}
		if weight == max {
			maxCount++
		}
	}

	restCakes := maxCakesCount - len(weights)
	if weight1 > weight2 { // 保证 weight1 是较小值
		weight1, weight2 = weight2, weight1
	}

	if weight1 > min || weight2 < max {
		return false
	}

	if weight1 == min {
		if weight2 == max {
			return true
		}
		return restCakes > 0
	} else {
		if weight2 == max {
			return restCakes > 0
		}
		return restCakes >= 2
	}
}
