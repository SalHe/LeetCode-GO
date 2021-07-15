package leetcode

// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
func combinationSum(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, i int)
	dfs = func(target, i int) {
		if i == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, i+1)
		if target-candidates[i] >= 0 {
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
