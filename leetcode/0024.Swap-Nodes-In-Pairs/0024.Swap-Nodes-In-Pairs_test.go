package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question24 struct {
	input  []int
	n      int
	output []int
}

func Test_swapPairs(t *testing.T) {
	qs := []question24{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{2, 1, 4, 3},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, q := range qs {
		// 构造输入链表
		var head *ListNode
		if len(q.input) > 0 {
			curr := &ListNode{Val: q.input[0]} // 头节点
			head = curr
			for i := 1; i < len(q.input); i++ {
				curr.Next = &ListNode{
					Val: q.input[i],
				}
				curr = curr.Next
			}
		}

		// 调用函数获取输出结果
		head = swapPairs(head)

		// 将输出链表转换为数组方便对比
		output := make([]int, 0)
		//head = head.Next
		for head != nil {
			output = append(output, head.Val)
			head = head.Next
		}

		// 对比结果
		assert.Equal(t, q.output, output)
	}
}
