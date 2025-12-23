package listnode

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// String 实现 Stringer 接口
func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}

	var sb strings.Builder
	current := n
	for current != nil {
		sb.WriteString(fmt.Sprintf("%d", current.Val))
		if current.Next != nil {
			sb.WriteString("->")
		}
		current = current.Next
	}
	return sb.String()
}
