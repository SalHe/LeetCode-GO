package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	dummy, curr := &ListNode{Next: head}, head
	pre, start := dummy, head
	var end, next *ListNode
	groupSize := 1
	for curr != nil {
		if groupSize == k {
			groupSize = 0
			end = curr
			next = curr.Next

			// 原链表: pre -> start -> ... -> end -> next
			// 翻转: start -> ... -> end
			// 得到：end -> ... -> start
			end.Next = nil
			reverse(start)
			pre.Next = end    // pre -> end -> ... -> start -> ...
			start.Next = next // pre -> end -> ... -> start -> next

			pre = start
			curr = next
			start = next
		} else {
			curr = curr.Next
		}
		groupSize++
	}
	return dummy.Next
}

func reverse(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}

	pre, cur := list, list.Next
	list.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return list
}
