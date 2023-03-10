package main

import (
	"fmt"
	"sort"
)

func solve(scores []int, x int, y int) int {
	sort.Ints(scores)

	// <= 分数线  ====> 淘汰
	// > 分数线  ====> 晋级

	for i := 0; i < len(scores); i++ {
		score := scores[i]
		for ; i < len(scores) && scores[i] == score; i++ {
		}
		i--

		qualified := len(scores) - 1 - i
		if x <= i+1 && i+1 <= y &&
			x <= qualified && qualified <= y {
			return score
		}
	}

	return -1
}

var _ = `
6 2 3
1 2 3 4 5 6

` // 3

var _ = `
100 45 60
169 218 217 369 493 93 460 551 754 431 34 932 644 576 560 12 324 176 867 29 399 62 848 980 549 98 938 121 923 661 744 767 873 419 764 317 164 656 997 399 278 552 437 853 993 5 332 643 759 541 803 919 293 849 562 335 396 983 635 424 807 537 596 640 255 868 425 37 186 482 568 571 863 505 987 339 764 890 745 518 497 242 133 22 122 556 502 301 576 917 605 732 601 366 955 302 996 654 560 898
` // 505

// 提交时间：2023-03-10 语言：Go 运行时间：299ms 占用内存：3924 状态：编译正确
func main() {

	var persons, x, y int
	fmt.Scanf("%d%d%d", &persons, &x, &y)
	scores := make([]int, persons)
	for i := 0; i < persons; i++ {
		fmt.Scanf("%d", &scores[i])
	}

	fmt.Println(solve(scores, x, y))
}
