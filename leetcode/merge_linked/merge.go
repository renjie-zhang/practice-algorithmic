package merge_linked

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var result = new(ListNode)
	result.Next = list1
	tempA := 0

	// find a
	cur := list1
	for cur != nil {
		tempA++
		if tempA == a {
			break
		}
		cur = cur.Next
	}
	// remove a->b
	// remove head node
	head := cur.Next
	cur.Next = list2
	// list2 tail
	for head != nil && tempA < b {
		tempA++
		head = head.Next
	}
	// find list2 tail
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head.Next
	return result.Next
}
