package leetcode

func subsets(nums []int) (result [][]int) {
	result = [][]int{{}} // 空集也是子集之一
	if len(nums) <= 0 {
		return
	}
	set := make([]int, len(nums))
	visited := make(map[int]struct{})
	solve(nums, 0, set, 0, 0, visited, &result)

	return result
}

func solve(nums []int, iNum int, set []int, iSet int, lSet int, visited map[int]struct{}, result *[][]int) {
	for ; iNum < len(nums); iNum++ {
		num := nums[iNum]
		if _, ok := visited[num]; ok {
			continue
		}
		visited[num] = struct{}{}
		set[iSet] = num
		lSet++
		*result = append(*result, append([]int{}, set[:lSet]...))
		solve(nums, iNum+1, set, iSet+1, lSet, visited, result)
		lSet--
		delete(visited, num)
	}
}
