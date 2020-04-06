package rotatelist

/***
链接： https://leetcode-cn.com/problems/rotate-list/
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}
	length := 1
	temp := head
	for temp.Next != nil {
		temp = temp.Next
		length++
	}
	if length >= k {
		for i := 0; i < k; i++ {
			temp := head
			for temp.Next.Next != nil {
				temp = temp.Next
			}
			TempNode := temp
			TempNode.Next.Next = head
			head = TempNode.Next
			TempNode.Next = nil
		}
	} else {
		mod := k % length
		for i := 0; i < mod; i++ {
			temp := head
			for temp.Next.Next != nil {
				temp = temp.Next
			}
			TempNode := temp
			TempNode.Next.Next = head
			head = TempNode.Next
			TempNode.Next = nil
		}
	}

	return head
}