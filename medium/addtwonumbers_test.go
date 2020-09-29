package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	param1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	param2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	expect := []int{7, 0, 8}
	result := addTwoNumbers(param1, param2)
	resultArray := listNode2Array(result)
	fmt.Println(resultArray)
	assert.Equal(t, expect, resultArray)
}
