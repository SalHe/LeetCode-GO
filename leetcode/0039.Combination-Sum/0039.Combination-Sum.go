package leetcode

// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
func combinationSum(candidates []int, target int) [][]int {
	// 先写出一个比较简单易懂的实现
	var result [][]int
	for _, candidate := range candidates {
		if candidate == target {
			result = append(result, []int{candidate})
		} else if candidate < target {
			mediaResults := combinationSum(candidates, target-candidate)
			for _, mediaResult := range mediaResults {
				if len(mediaResult) <= 0 || candidate > mediaResult[0] {
					continue
				}
				r := make([]int, len(mediaResult)+1)
				r[0] = candidate
				copy(r[1:], mediaResult)
				result = append(result, r)
			}
		}
	}
	return result
}
