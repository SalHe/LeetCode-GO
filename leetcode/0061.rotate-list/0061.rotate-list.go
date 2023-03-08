package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	head = &ListNode{Next: head}

	fast := head
	// 1. n > k，那么还没遍历完链表就定位出fast了
	// 2. n < k，那么等遍历完链表的时候循环还未结束，但是现在实际上我们可以将遍历次数缩小（k%=i）
	// 3. n == k 或者说 n可以由k整除的话，直接退出
	for i := 0; i < k; i++ {
		fast = fast.Next

		// 提前到达链尾
		if fast == nil {
			k %= i // 这时候i肯定是链表长度

			if k == 0 {
				// 发现链表长度刚好是k的整数倍，那就没有必要循环了
				return head.Next
			}

			// 为了防止循环提前终止，这里要给他加上一个n（因为此时k<i就会跳出循环了，不是我们想要的）
			k += i
			fast = head.Next
		} else if fast.Next == nil && i == k-1 {
			// 这里是针对k==n的情况
			// n==k时循环不会到链尾就结束了，就不知道他们n==k
			return head.Next
		}
	}

	// 可以保证slow定位到倒数k%n个元素
	var preSlow, preFast *ListNode
	slow := head
	for fast != nil {
		preFast = fast
		preSlow = slow
		fast = fast.Next
		slow = slow.Next
	}

	head.Next, preFast.Next, preSlow.Next = slow, head.Next, nil
	return head.Next
}
