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
package find_first_and_last_position

func SearchRange(nums []int, target int) []int {
	var targetValue = []int{-1, -1}
	// 把第一出现的地方叫做下边界 最后一次出现的地方为上边界
	// 上边界的条件：
	// 1 该值右边一个数不与之相等
	// 2 该书左边没有数，也就是这个数是最后一个

	a := searchLowerBound(nums, target, 0, len(nums)-1)
	targetValue[0] = a
	b := searchUpperBound(nums, target, 0, len(nums)-1)
	targetValue[1] = b
	// 下边界的条件：
	// 1 该值左边一个数不与之相等
	// 2 该书左边没有数，也就是这个数是第一个
	return targetValue
}

func searchLowerBound(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	midlle := low + (high-low)/2
	// 查看是否为第一个，或者左边的数比她小 此时下边界成立
	if nums[midlle] == target && (midlle == 0 || nums[midlle-1] < target) {
		return midlle
	}
	if target <= nums[midlle] {
		return searchLowerBound(nums, target, low, midlle-1)
	} else {
		return searchLowerBound(nums, target, midlle+1, high)
	}
}

func searchUpperBound(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	midlle := low + (high-low)/2
	// 查看是否为第一个，或者左边的数比她小 此时下边界成立
	if nums[midlle] == target && (midlle == len(nums)-1 || nums[midlle+1] > target) {
		return midlle
	}
	if target < nums[midlle] {
		return searchUpperBound(nums, target, low, midlle-1)
	} else {
		return searchUpperBound(nums, target, midlle+1, high)
	}
}
