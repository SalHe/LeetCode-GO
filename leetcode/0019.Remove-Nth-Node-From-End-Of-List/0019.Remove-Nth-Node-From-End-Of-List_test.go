package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question19 struct {
	input  []int
	n      int
	output []int
}

func Test_removeNthFromEndWithoutHead(t *testing.T) {
	qs := []question19{
		{
			input:  []int{1, 2},
			n:      2,
			output: []int{2},
		},
		{
			input:  []int{1},
			n:      1,
			output: []int{},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			n:      2,
			output: []int{1, 2, 3, 5},
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
		head = removeNthFromEnd(head, q.n)

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

func Test_removeNthFromEndWithHead(t *testing.T) {
	qs := []question19{
		{
			input:  []int{1, 2},
			n:      2,
			output: []int{2},
		},
		{
			input:  []int{1},
			n:      1,
			output: []int{},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			n:      2,
			output: []int{1, 2, 3, 5},
		},
	}

	for _, q := range qs {
		// 构造输入链表
		var head *ListNode
		if len(q.input) > 0 {
			curr := &ListNode{} // 头节点
			head = curr
			for i := 0; i < len(q.input); i++ {
				curr.Next = &ListNode{
					Val: q.input[i],
				}
				curr = curr.Next
			}
		}

		// 调用函数获取输出结果
		head = removeNthFromEnd(head, q.n)

		// 将输出链表转换为数组方便对比
		output := make([]int, 0)
		head = head.Next
		for head != nil {
			output = append(output, head.Val)
			head = head.Next
		}

		// 对比结果
		assert.Equal(t, q.output, output)
	}
}
