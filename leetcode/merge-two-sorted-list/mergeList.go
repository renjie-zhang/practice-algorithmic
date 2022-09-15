/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package merge_two_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := head
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}
	for l1 != nil || l2 != nil {
		if (l1 != nil && l2 != nil) && l1.Val > l2.Val {
			t2 := l2
			t2.Next = nil
			l2 = l2.Next
			head.Next = t2
			head = head.Next

			t1 := l1
			t1.Next = nil
			l1 = l1.Next
			head.Next = t1
			head = head.Next

		} else if (l1 != nil && l2 != nil) && l1.Val < l2.Val {
			t1 := l1
			t1.Next = nil
			l1 = l1.Next
			head.Next = t1
			head = head.Next

			t2 := l2
			t2.Next = nil
			l2 = l2.Next
			head.Next = t2
			head = head.Next
		} else if (l1 != nil && l2 != nil) && l1.Val == l2.Val {
			t1 := l1
			t1.Next = nil
			l1 = l1.Next
			t2 := l2
			l2 = l2.Next
			t2.Next = nil

			head.Next = t1
			head = head.Next

			head.Next = t2
			head = head.Next
		} else if l1 == nil {
			t2 := l2
			t2.Next = nil
			l2 = l2.Next
			head.Next = t2
			head = head.Next
		} else {
			t1 := l1
			t1.Next = nil
			l1 = l1.Next
			head.Next = t1
			head = head.Next
		}

	}
	return temp.Next
}
