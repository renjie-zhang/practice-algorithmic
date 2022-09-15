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

import "testing"

func TestMergeTwoLists(t *testing.T) {
	one1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	two1 := &ListNode{
		Val:  2,
		Next: nil,
	}

	four1 := &ListNode{
		Val:  4,
		Next: nil,
	}
	first := &ListNode{
		Val:  0,
		Next: nil,
	}

	second := &ListNode{
		Val:  0,
		Next: nil,
	}
	three2 := &ListNode{
		Val:  3,
		Next: nil,
	}
	one2 := &ListNode{
		Val:  1,
		Next: nil,
	}

	four2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	temp1 := first
	first.Next = one1
	first = first.Next
	first.Next = two1
	first = first.Next
	first.Next = four1

	temp2 := second
	second.Next = one2
	second = second.Next
	second.Next = three2
	second = second.Next
	second.Next = four2

	MergeTwoLists(temp1.Next, temp2.Next)

}
