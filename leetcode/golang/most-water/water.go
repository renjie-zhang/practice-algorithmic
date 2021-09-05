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
package most_water

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	var max float64
	j := len(height) -1
	for i := 0; i < j;{
		var minHeight int
		if height[i] < height[j]{
			minHeight= height[i]
			i++
		}else {
			minHeight = height[j]
			j--
		}
		max = math.Max(max,float64((j-i+1)*minHeight))
	}
	return int(max)
}

func Max(){
	var te = []int{1,8,6,2,5,4,8,3,7}
	fmt.Printf("Value is %d",maxArea(te))
}
