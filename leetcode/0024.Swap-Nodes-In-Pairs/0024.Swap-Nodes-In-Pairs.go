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

func swapPairs(head *ListNode) *ListNode {
	if head == nil { // 空链表
		return nil
	}
	header := &ListNode{Next: head}
	pre := header
	for pre != nil && pre.Next != nil {
		node1, node2 := pre.Next, pre.Next.Next
		if node2 == nil { // 奇数个结点
			break
		}

		// ... -> pre -> node1 -> node2 -> next -> ...
		// ... -> pre -> node2 -> node1 -> next -> ...
		pre.Next = node2        // ... -> pre -> node2 -> next -> ...,  node1 -> node2 -> next -> ...
		node1.Next = node2.Next // ... -> pre -> node2 -> next -> ...,  node1 -> next -> ...
		node2.Next = node1      // ... -> pre -> node2 -> node1 -> next -> ...

		pre = node1
	}
	return header.Next
}
