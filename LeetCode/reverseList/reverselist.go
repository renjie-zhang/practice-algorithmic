package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 定义一个辅助结点
	current := head
	// 记录下一个结点的位置
	var next *ListNode
	// 创建一个反转链表的头结点
	reverselist := &ListNode{Val: 0}
	for current != nil {
		//保存当前结点的下一个结点
		next = current.Next
		// 将current的下一个结点指向新链表的前端
		current.Next = reverselist.Next
		// 将反转链表的结点指向当前结点
		reverselist.Next = current
		// 完成之后current后移
		current = next
	}
	return reverselist.Next
}
