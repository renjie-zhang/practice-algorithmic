/*
 * @Descripttion:删除链表的倒数第n个结点
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-10-24 17:30:08
 * @LastEditTime: 2019-11-05 21:12:15
 */
package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
删除链表的倒数第n个结点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	temp := head
	for temp.Next != nil {
		length = length + 1
		temp = temp.Next
	}
	if n < 0 || n > length {
		return nil
	}
	if n == length {
		return head
	}

	current := head

	for i := length - n - 1; i > 0; i-- {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}
