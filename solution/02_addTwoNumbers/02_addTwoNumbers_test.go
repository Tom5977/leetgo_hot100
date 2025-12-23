package addtwonumbers

import (
	"fmt"
	"leetgo_hot100/utils/listnode"
	"testing"
)

func Test02(t *testing.T) {
	l1 := listnode.NewFromSlice([]int{2, 4, 3})
	l2 := listnode.NewFromSlice([]int{5, 6, 4})
	fmt.Print(addTwoNumbers(l1, l2))
}
