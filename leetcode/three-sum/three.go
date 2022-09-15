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
package three_sum

import "sort"

// 使用双指针+排序
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	k := 0
	for ; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for i < j {
			temp := nums[i] + nums[j] + nums[k]
			if temp < 0 {
				i += 1
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
			} else if temp > 0 {
				j -= 1
				for i < j && nums[j] == nums[j+1] {
					j -= 1
				}
			} else {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				i += 1
				j -= 1
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
				for i < j && nums[j] == nums[j+1] {
					j -= 1
				}
			}
		}
	}
	return result
}
