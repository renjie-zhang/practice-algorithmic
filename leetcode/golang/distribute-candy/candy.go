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
package distribute_candy

import "container/list"

func DistributeCandies(candyType []int) int {
	var s list.List
	var r  = make(map[int]bool,len(candyType))
	for i := 0; i < len(candyType); i++ {
		if _,ok := r[candyType[i]];!ok{
			s.PushBack(candyType[i])
			r[candyType[i]] = true
		}
	}
	var length = s.Len()
	var max = len(candyType)/2
	if length <= max{
		return length
	}
	return max
}



