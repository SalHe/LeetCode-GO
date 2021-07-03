package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 以无头结点的方式处理，并利用现有结点
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) > 2 {
		mid := len(lists) / 2
		left, right := mergeKLists(lists[:mid]), mergeKLists(lists[mid:])
		return mergeKLists([]*ListNode{left, right})
	} else {

		cur, allNil := make([]*ListNode, len(lists)), true
		for i := range lists {
			if lists[i] != nil {
				allNil = false
			}
			cur[i] = lists[i]
		}
		if allNil {
			return nil
		}

		newList := &ListNode{}
		curNode := newList

		for {
			// 寻找包含最小值的列表的当前节点
			var minNode *ListNode = nil
			iList := -1
			for i := len(cur) - 1; i >= 0; i-- {
				if cur[i] == nil {
					continue
				}
				if minNode == nil || cur[i].Val < minNode.Val {
					minNode = cur[i]
					iList = i
				}
			}

			// 找不到这样的结点，说明合并完了
			if iList < 0 {
				break
			}

			cur[iList] = cur[iList].Next

			curNode.Next = minNode
			curNode = minNode
		}
		return newList.Next
	}
}
