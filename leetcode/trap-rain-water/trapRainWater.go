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
package trap_rain_water

func Trap(height []int) int {
	length := len(height)
	var answer int
	leftMax, rightMax := 0, 0
	left, right := 0, length-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				answer += leftMax - height[left]
			}
			left += 1
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				answer += rightMax - height[right]
			}
			right -= 1
		}

	}

	return answer
}
