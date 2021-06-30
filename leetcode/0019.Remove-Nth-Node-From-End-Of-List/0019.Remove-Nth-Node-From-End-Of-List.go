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

// 这个题目就是以不带头节点来做的。。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 假设给的是不带头的
	head = &ListNode{Next: head}

	// 以带头节点的方式进行处理
	cache, pCache := make([]*ListNode, n+1), 1
	cache[pCache] = head
	curr := head
	for curr != nil {
		// 将遍历过的节点的最后n个放至缓存区
		// 最后会用到
		cache[pCache] = curr
		pCache = (pCache + 1) % (n + 1)
		curr = curr.Next
	}
	target := cache[pCache]
	if next := target.Next; next != nil {
		target.Next = next.Next
	}
	return head.Next
}
