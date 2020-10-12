package medium

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Arr := listNode2Array(l1)
	l2Arr := listNode2Array(l2)
	maxLen := getMaxLen(l1Arr, l2Arr)
	var newArr []int
	for index := 0; index < maxLen; index++ {
		l1Item := getArrItem(l1Arr, index)
		l2Item := getArrItem(l2Arr, index)
		newArr = append(newArr, l1Item+l2Item)
	}

	for index, item := range newArr {
		if item > 9 {
			if len(newArr) == index+1 {
				newArr[index] = item % 10
				newArr = append(newArr, 1)
			} else {
				newArr[index] = item % 10
				newArr[index+1] += 1
			}
		}
	}

	return intArray2ListNode(newArr)
}

func getArrItem(arr []int, index int) int {
	if len(arr) < index+1 {
		return 0
	} else {
		return arr[index]
	}
}

func getMaxLen(l1Arr, l2Arr []int) int {
	maxLen := len(l1Arr)
	if len(l2Arr) > maxLen {
		maxLen = len(l2Arr)
	}
	return maxLen
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
