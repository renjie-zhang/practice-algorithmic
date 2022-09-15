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
package longest_substring_without_repeat

import (
	"math"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	tempMap := make(map[string]int)
	max := 0
	tempS := strings.Split(s, "")
	for i, j := 0, 0; j < len(tempS); j++ {
		tempString := tempS[j]
		if _, ok := tempMap[tempString]; ok {
			i = int(math.Max(float64(i), float64(tempMap[tempString]+1)))
		}
		tempMap[tempS[j]] = j
		max = int(math.Max(float64(max), float64(j-i+1)))
	}
	return max
}
