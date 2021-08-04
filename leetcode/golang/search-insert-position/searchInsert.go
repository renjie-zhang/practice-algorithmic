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
package search_insert_position

func SearchInsert(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	answer := length
	for left < right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return answer
}
