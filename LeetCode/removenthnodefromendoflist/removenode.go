package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length = 1
	temp := head
	for temp.Next != nil {
		length++
		temp = temp.Next
	}
	if n < 0 || n > length {
		return nil
	}
	if n == length {
		head = head.Next
		return head
	}
	current := head
	for i := length - n - 1; i > 0; i-- {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}
