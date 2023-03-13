package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	head = &ListNode{Next: head}
	pre, cur := head, head.Next
	for cur != nil {
		start := cur
		for cur != nil && start.Val == cur.Val {
			cur = cur.Next
		}
		if start.Next != cur {
			pre.Next = cur
		} else {
			pre = pre.Next
		}
	}

	return head.Next
}
