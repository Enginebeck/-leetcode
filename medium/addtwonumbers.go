package medium

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// todo 1. convert to array
	//      2. convert to integer
	//      3. sum
	//      4. split and reverse
	l1RevArr := listNode2Array(l1)
	l2RevArr := listNode2Array(l2)
	maxLen := len(l1RevArr)
	if len(l2RevArr) > maxLen {
		maxLen = len(l2RevArr)
	}
	l1number := getNumber(l1RevArr, maxLen)
	l2number := getNumber(l2RevArr, maxLen)
	sumResult := l1number + l2number
	fmt.Println(sumResult)

	return getResult(sumResult, maxLen)
}

func getResult(sumResult int, maxLen int) *ListNode {
	var result *ListNode
	for index := 0; index < maxLen; {
		result = &ListNode{sumResult / int(math.Pow10(maxLen-1-index)), result}
		sumResult = sumResult % int(math.Pow10(maxLen-1-index))
		index++
	}
	return result
}

func getNumber(arr []int, maxLen int) int {
	result := 0
	for index := 0; index < maxLen; {
		result += arr[index] * int(math.Pow10(maxLen-index-1))
		index++
	}
	return result
}

func listNode2Array(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
