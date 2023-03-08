package leetcode

type ListNode struct {
	headed bool
	Val    int
	Next   *ListNode
}

func BuildHeadedList(numbers []int) *ListNode {
	return &ListNode{headed: true, Next: BuildList(numbers)}
}

func BuildList(numbers []int) *ListNode {
	var head *ListNode
	if len(numbers) > 0 {
		curr := &ListNode{Val: numbers[0]} // 头节点
		head = curr
		for i := 1; i < len(numbers); i++ {
			curr.Next = &ListNode{
				Val: numbers[i],
			}
			curr = curr.Next
		}
	}
	return head
}

func (n *ListNode) ToArray() (output []int) {
	head := n
	if n.headed {
		head = head.Next
	}
	for head != nil {
		output = append(output, head.Val)
		head = head.Next
	}
	return
}
