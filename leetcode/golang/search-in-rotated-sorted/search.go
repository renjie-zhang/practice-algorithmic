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
package search_in_rotated_sorted

func Search(nums []int, target int) int {
	return binarysearch(nums, target, 0, len(nums)-1)
}

func binarysearch(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	middle := low + (high-low)/2
	// 如果nums[left] < nums[middles]
	if nums[middle] == target {
		return middle
	}
	if nums[low] <= nums[middle] {
		if nums[low] <= target && target < nums[middle] {
			return binarysearch(nums, target, low, middle-1)
		}
		return binarysearch(nums, target, middle+1, high)
	} else {
		if nums[middle] < target && target <= nums[high] {
			return binarysearch(nums, target, middle+1, high)
		}
		return binarysearch(nums, target, low, middle-1)
	}
}
