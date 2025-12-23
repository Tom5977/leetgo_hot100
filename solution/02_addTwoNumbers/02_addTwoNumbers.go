package addtwonumbers

import "leetgo_hot100/utils/listnode"

type ListNode = listnode.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := 0
	p := dummy
	for l1 != nil || l2 != nil || temp > 0 {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + temp
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		temp = sum / 10

	}
	return dummy.Next
}
