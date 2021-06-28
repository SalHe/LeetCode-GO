package leetcode

var lettersTable = []string{
	"", // 0
	"", // 1
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	length := len(digits)
	nums := make([]int32, length)
	allPossible := 1
	for i, c := range digits {
		nums[i] = c - '0'
		allPossible *= len(lettersTable[nums[i]])
	}

	// 可以使用递归实现，不过这里不使用递归
	combinations := make([]string, allPossible)
	buff := make([]uint8, length)
	letterIndex := make([]int, length)
	currCom, depth := 0, 0
	for depth >= 0 { // 递归出口
		if letterIndex[depth] >= len(lettersTable[nums[depth]]) { // 对应层次的已经递归完毕
			depth--
			continue
		}

		buff[depth] = lettersTable[nums[depth]][letterIndex[depth]]
		letterIndex[depth]++
		if depth < length-1 {
			depth++
			letterIndex[depth] = 0
		} else if depth == length-1 {
			combinations[currCom] = string(buff)
			currCom++
		}
	}
	return combinations
}
