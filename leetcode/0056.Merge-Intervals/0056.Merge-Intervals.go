package leetcode

import "sort"

type _intervals struct {
	target [][]int
}

func (a _intervals) Len() int {
	return len(a.target)
}

func (a _intervals) Less(i, j int) bool {
	return a.target[i][0] < a.target[j][0]
}

func (a _intervals) Swap(i, j int) {
	a.target[i], a.target[j] = a.target[j], a.target[i]
}

func sortIntervals(intervals [][]int) {
	a := _intervals{target: intervals}
	sort.Sort(a)
}

func merge(intervals [][]int) (result [][]int) {
	sortIntervals(intervals)
	i := 0
	for i < len(intervals) {
		j := i + 1
		// [a,b] ... [c,d] (c<=b都可以合并)
		end := intervals[i][1]
		for j < len(intervals) && intervals[j][0] <= end {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		result = append(result, []int{intervals[i][0], end})
		i = j
	}
	return
}
