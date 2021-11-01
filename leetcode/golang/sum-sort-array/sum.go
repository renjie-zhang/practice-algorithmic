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
package sum_sort_array

func GetSumAbsoluteDifferences(nums []int) []int {
	var result []int
	var length = len(nums)
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	result = append(result,sum-nums[0]*length)
	for i := 1; i < length; i++ {
		var temp = nums[i] - nums[i-1]
		result = append(result,result[i-1]-(length-i*2)*temp)
	}

	return result
}
