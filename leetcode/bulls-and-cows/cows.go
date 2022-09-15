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
package bulls_and_cows

import "strconv"

func GetHint(secret string, guess string) string {
	var cowsTemp [10]int
	var bullsTemp [10]int
	var bullsCount int
	var cowsCount int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i]{
			bullsCount ++
		}else {
			cowsTemp[secret[i]-'0']++
			bullsTemp[guess[i] - '0']++
		}
	}

	for i := 0; i < 10; i++ {
		cowsCount += getMin(bullsTemp[i],cowsTemp[i])
	}

	return strconv.Itoa(bullsCount)+"A"+strconv.Itoa(cowsCount)+"B"
}
func getMin(x,y int) int{
	if x > y{
		return y
	}
	return x
}
