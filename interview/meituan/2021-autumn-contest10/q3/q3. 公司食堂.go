package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 提交时间：2023-03-10 语言：Go 运行时间：4001ms 占用内存：5604 状态：编译错误(实际上是未完全通过测试用例，Q4同)
func main() {
	sb := &strings.Builder{}

	var cases int
	fmt.Scanf("%d ", &cases)
	for i := 0; i < cases; i++ {

		var tables int
		fmt.Scanf("%d ", &tables)
		sb.Grow(tables * 2)

		seats := make([][]int, 2)
		for j := 0; j < 2; j++ {
			seats[j] = make([]int, tables)
			for k := 0; k < tables; k++ {
				seats[j][k] = math.MaxInt
			}
		}
		seatsCount := make([]int, 2)
		var s string
		fmt.Scanf("%s ", &s)
		for j, b := range s {
			if b == '2' {
				continue
			}
			seats[b-'0'][seatsCount[b-'0']] = j + 1
			seatsCount[b-'0']++
		}

		var personsCount int
		fmt.Scanf("%d ", &personsCount)
		isMale := make([]bool, personsCount)
		fmt.Scanf("%s ", &s)
		for j, gender := range s {
			isMale[j] = gender == 'M'
		}

		result := solve(seats, seatsCount, isMale)

		for _, r := range result {
			sb.WriteString(strconv.Itoa(r))
			sb.WriteString("\n")
		}
	}

	fmt.Println(sb.String())
}

func solve(seats [][]int, count []int, male []bool) (result []int) {
	result = make([]int, len(male))
	take := func(x int) int {
		target := 0
		for i := 1; i < len(seats[x]); i++ {
			if seats[x][i] != math.MaxInt && seats[x][i] < seats[x][target] {
				target = i
			}
		}
		r := seats[x][target]
		// seats[x] = append(seats[x][:target], seats[x][target+1:]...)
		seats[x][target] = math.MaxInt
		count[x]--
		if x == 0 {
			// seats[x+1] = append(seats[x+1], r)
			for i := 0; i < len(seats[x+1]); i++ {
				if seats[x+1][i] == math.MaxInt {
					seats[x+1][i] = r
					count[x+1]++
					break
				}
			}
		}
		return r
	}

	p := 0
	for _, isMale := range male {
		if isMale {
			if count[1] > 0 {
				result[p] = take(1)
			} else if count[0] > 0 { // 按理来说，只要给定输入没问题，那么这里其实不需要判断
				result[p] = take(0)
			}
		} else {
			if count[0] > 0 {
				result[p] = take(0)
			} else if count[1] > 0 {
				result[p] = take(1)
			}
		}
		p++
	}
	return
}
