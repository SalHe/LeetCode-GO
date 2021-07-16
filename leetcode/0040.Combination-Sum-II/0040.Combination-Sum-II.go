package leetcode

import "sort"

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var comb []int
	var dfs func(target, i int)
	dfs = func(target, i int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		if i >= len(candidates) {
			return
		}
		for d := 1; d+i < len(candidates); d++ {
			if candidates[i+d] != candidates[i] {
				dfs(target, i+d)
				break
			}
		}

		if target-candidates[i] >= 0 {
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
