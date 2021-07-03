package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question25 struct {
	input  []int
	k      int
	output []int
}

func Test_reverseKGroup(t *testing.T) {
	qs := []question25{
		//{
		//	input:  []int{1, 2, 3, 4, 5},
		//	k:      2,
		//	output: []int{2, 1, 4, 3, 5},
		//},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      3,
			output: []int{3, 2, 1, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      1,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1},
			k:      1,
			output: []int{1},
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
		head = reverseKGroup(head, q.k)

		// 将输出链表转换为数组方便对比
		output := make([]int, 0)
		for head != nil {
			output = append(output, head.Val)
			head = head.Next
		}

		// 对比结果
		assert.Equal(t, q.output, output)
	}
}
