package medium

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1rev := reverse(l1)
	l2rev := reverse(l2)
	result := &ListNode{0, nil}
	for {
		if result.Val != 0 {
			result.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			result = result.Next
		}
		if l1rev != nil {
			result.Val += l1rev.Val
			l1rev = l1rev.Next
		}
		if l2rev != nil {
			result.Val += l2rev.Val
			l2rev = l2rev.Next
		}
		if l2rev == nil && l1rev == nil {
			break
		}
	}
	return reverse(result)
}

func reverse(node *ListNode) *ListNode {
	var listNode *ListNode
	for {
		temp := listNode
		listNode = node
		listNode.Next = temp
		if node.Next == nil {
			break
		} else {
			node = node.Next
		}
	}
	return listNode
}
