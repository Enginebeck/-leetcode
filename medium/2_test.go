package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		Param1 *ListNode
		Param2 *ListNode
		Expect *ListNode
	}{{
		intArray2ListNode([]int{2, 4, 3}),
		intArray2ListNode([]int{5, 6, 4}),
		intArray2ListNode([]int{7, 0, 8}),
	}, {
		intArray2ListNode([]int{5}),
		intArray2ListNode([]int{5}),
		intArray2ListNode([]int{0, 1}),
	}, {
		intArray2ListNode([]int{1, 8}),
		intArray2ListNode([]int{0}),
		intArray2ListNode([]int{1, 8}),
	}, {
		intArray2ListNode([]int{3, 6, 9, 8}),
		intArray2ListNode([]int{2, 3, 6, 4}),
		intArray2ListNode([]int{5, 9, 5, 3, 1}),
	}}
	for _, caseItem := range cases {
		result := addTwoNumbers(caseItem.Param1, caseItem.Param2)
		resultArray := listNode2Array(result)
		expectArr := listNode2Array(caseItem.Expect)
		assert.Equal(t, expectArr, resultArray)
	}
}
