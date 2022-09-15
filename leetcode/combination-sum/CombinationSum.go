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
package combination_sum

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	dfs(candidates, nil, target, 0, &result)
	return result
}

// 减法的方式进行解决
func dfs(candidates []int, nums []int, target int, left int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	for i := left; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, result)
	}
}
