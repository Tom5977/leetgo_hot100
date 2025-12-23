package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "example_1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example_2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example_3_with_negative",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "no_solution", // 测试边界或异常情况
			nums:     []int{1, 2, 3},
			target:   7,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
