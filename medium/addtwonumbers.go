package medium

import (
	"fmt"
	"math"
	"strconv"
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
	l1Arr := listNode2Array(l1)
	l2Arr := listNode2Array(l2)
	maxLen := len(l1Arr)
	if len(l2Arr) > maxLen {
		maxLen = len(l2Arr)
	}
	l1number := getNumber(l1Arr, maxLen)
	l2number := getNumber(l2Arr, maxLen)
	sumResult := l1number + l2number
	fmt.Println(l1Arr, l2Arr)
	fmt.Println(l1number, l2number)
	return getResult(sumResult)
}

func getResult(sumResult int) *ListNode {
	var result *ListNode
	maxLen := len(strconv.Itoa(sumResult))
	for index := 0; index < maxLen; {
		result = &ListNode{sumResult / int(math.Pow10(maxLen-1-index)), result}
		sumResult = sumResult % int(math.Pow10(maxLen-1-index))
		index++
	}
	return result
}

func getNumber(arr []int, maxLen int) int {
	result := 0
	for index := maxLen - 1; index >= 0; {
		if index >= len(arr) {
			index = len(arr) - 1
		}
		result += arr[index] * int(math.Pow10(index))
		index--
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

func intArray2ListNode(arr []int) *ListNode {
	var result *ListNode
	for index := len(arr) - 1; index >= 0; {
		result = &ListNode{
			Val:  arr[index],
			Next: result,
		}
		index--
	}
	return result
}
