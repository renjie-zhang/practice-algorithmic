package merge_linked

import (
	"fmt"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	var temp1 = []int{0, 1, 2, 3, 4, 5}
	var list1 = new(ListNode)
	head := list1
	for i := 0; i < len(temp1); i++ {
		t := new(ListNode)
		t.Val = temp1[i]
		head.Next = t
		head = head.Next
	}

	var temp2 = []int{1000000, 1000001, 1000002}
	var list2 = new(ListNode)
	head2 := list2
	for i := 0; i < len(temp2); i++ {
		t := new(ListNode)
		t.Val = temp2[i]
		head2.Next = t
		head2 = head2.Next
	}
	between := MergeInBetween(list1.Next, 3, 4, list2.Next)
	for between != nil {
		fmt.Print(between.Val)
		fmt.Print("\t")
		between = between.Next
	}
}
