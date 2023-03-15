// Package leetcode
// @Author: SalHe
// @Date:   2023/3/15 11:45
package leetcode

func partition(head *ListNode, x int) *ListNode {
	l1, l2 := new(ListNode), new(ListNode)
	p1, p2 := l1, l2

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			p1.Next = &ListNode{Val: cur.Val}
			p1 = p1.Next
		} else {
			p2.Next = &ListNode{Val: cur.Val}
			p2 = p2.Next
		}
	}

	p1.Next = l2.Next

	return l1.Next
}

func partitionNoNew(head *ListNode, x int) *ListNode {
	head = &ListNode{Next: head}
	cur, pre := head.Next, head
	smaller := head
	for cur != nil {
		if cur.Val < x {
			if smaller.Next != cur {
				smaller.Next, cur.Next, pre.Next = cur, smaller.Next, cur.Next
			} else {
				pre = smaller.Next
			}
			cur, smaller = pre.Next, cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head.Next
}
