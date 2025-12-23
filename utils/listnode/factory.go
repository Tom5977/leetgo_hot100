package listnode

func NewFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	curr := dummy

	for _, val := range vals {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}
