package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question23 struct {
	lists   [][]int
	newList []int
}

func Test_mergeKLists(t *testing.T) {
	qs := []question23{
		{
			lists: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			newList: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			lists:   [][]int{},
			newList: []int{},
		},
		{
			lists:   [][]int{{}},
			newList: []int{},
		},
	}

	for _, q := range qs {
		lists, length := make([]*ListNode, len(q.lists)), 0
		for i := range lists {
			cur := lists[i]
			length += len(q.lists[i])
			for j := range q.lists[i] {
				if cur == nil {
					cur = &ListNode{Val: q.lists[i][j]}
					lists[i] = cur
				} else {
					cur.Next = &ListNode{Val: q.lists[i][j]}
					cur = cur.Next
				}
			}
		}

		newList := mergeKLists(lists)
		cur, arr, i := newList, make([]int, length), 0
		for cur != nil {
			arr[i] = cur.Val
			i++
			cur = cur.Next
		}

		assert.ElementsMatch(t, q.newList, arr)

	}
}
