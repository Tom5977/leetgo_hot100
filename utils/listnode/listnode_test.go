package listnode

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	nums := []int{3, 6, 9}
	l1 := NewFromSlice(nums)
	fmt.Print(l1)
}
