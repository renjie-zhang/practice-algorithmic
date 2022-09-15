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
package range_addition_ii

func MaxCount(m int, n int, ops [][]int) int {
	min1,min2 := m,n
	for _, op := range ops {
		min1 = getMin(min1,op[0])
		min2 = getMin(min2,op[1])
	}
	return min1*min2
}

func getMin(x,y int) int{
	if x > y{
		return y
	}
	return x
}
