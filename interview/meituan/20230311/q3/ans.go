package main

import (
	"fmt"
	"sort"
)

// 9%

func main() {
	var stars int
	fmt.Scanf("%d\n", &stars)
	ranges := make([]rg, stars)
	for i := 0; i < stars; i++ {
		fmt.Scanf("%d", &ranges[i].from)
		ranges[i].stars = 1
	}
	fmt.Scanln()
	for i := 0; i < stars; i++ {
		fmt.Scanf("%d", &ranges[i].to)
	}

	a, b := solve(ranges)
	fmt.Printf("%d %d\n", a, b)
}

type rg struct {
	from  int
	to    int
	stars int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(rgs []rg) (int, int) {
	sort.Slice(rgs, func(i, j int) bool {
		return rgs[i].from < rgs[j].from
	})

	// var sames []rg
	for merged := true; merged; {
		merged = false
		for i := 0; i < len(rgs)-1; i++ {
			if rgs[i].to >= rgs[i+1].from {
				// rgs[i].from = max(rgs[i].from, rgs[i+1].from)
				// rgs[i].to = min(rgs[i].to, rgs[i+1].to)
				rgs[i].from = rgs[i+1].from
				rgs[i].stars++
				if rgs[i].to >= rgs[i+1].to {
					rgs[i].to = rgs[i+1].to
					rgs = append(rgs[:i+1], rgs[i+2:]...)
				} else {
					for j := i + 2; j < len(rgs); j++ {
						if rgs[i+1].from >= rgs[j].to {
							rgs[i+1].from = rgs[j].to + 1
						}
					}
					// rgs[i+1].from++
				}
			}
		}
	}

	times := make(map[int]struct{})
	maxStars := 0
	for _, r := range rgs {
		if r.stars > maxStars {
			maxStars = r.stars
		}
	}
	for _, r := range rgs {
		if r.stars != maxStars {
			continue
		}
		for i := r.from; i <= r.to; i++ {
			times[i] = struct{}{}
		}
	}

	return maxStars, len(times)
}

// 仅用于方便复制测试用例，与代码无关
var _ = `
3
2 1 5
6 3 7

`
