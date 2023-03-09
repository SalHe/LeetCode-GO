package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	size := len(matrix[0])
	var i int
	for i = 0; i < rows && target > matrix[i][0]; i++ {
	}

	if i == 0 {
		return target == matrix[0][0]
	}
	if i == rows && target > matrix[rows-1][size-1] {
		return false
	}
	if i < rows && target == matrix[i][0] {
		return true
	}
	m := matrix[i-1]
	left, right := 0, size-1
	for left < right {
		mid := (right-left+1)/2 + left
		if m[mid] == target {
			return true
		} else if m[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return false
}
