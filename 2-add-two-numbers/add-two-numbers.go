/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var overflow int = 0
	var startNode *ListNode = &ListNode{}
	var currentNode *ListNode = startNode
	for l1 != nil || l2 != nil || overflow != 0 {
		result := getNodeValue(l1) + getNodeValue(l2) + overflow
    
		if result >= 10 {
			overflow = 1
			result = result % 10
		} else {
			overflow = 0
		}
		currentNode.Val = result

		l1 = moveToNextNode(l1)
		l2 = moveToNextNode(l2)
		if l1 != nil || l2 != nil || overflow != 0 {
			newNode := &ListNode{}
			currentNode.Next = newNode
			currentNode = newNode
		}
	}

	return startNode
}

func moveToNextNode(l *ListNode) *ListNode {
	if l != nil {
		l = l.Next
	}
	return l
}

func getNodeValue(l *ListNode) int {
	if l != nil {
		return l.Val
	}
	return 0
}