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
package relative_sort_array

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	tempmap := make(map[int]int)
	for k,v := range arr2{
		tempmap[v] = k
	}
	sort.Slice(arr1, func(i, j int) bool {
		x,y := arr1[i],arr1[j]
		vx,okx := tempmap[x]
		vy,oky := tempmap[y]
		if okx&&oky{
			return vx < vy
		}
		if oky || okx{
			return okx
		}
		return x < y
	});
	return arr1
}
