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


/**
In this Kata, you will be given a number n (n > 0) and your task will be to return the smallest square number N (N > 0)
such that n + N is also a perfect square. If there is no answer, return -1 (nil in Clojure, Nothing in Haskell).
More examples in test cases.
Good luck.
Please don't forget to rate this kata.
https://www.codewars.com/kata/5edc8c53d7cede0032eb6029/train/go
 */
package simple_square_numbers

import "math"

func Solve(n int) int {

	for b := int(math.Sqrt(float64(n))) +1; getSquare(b)-getSquare(b-1) <= n ;b++  {
		temp := math.Sqrt(float64(getSquare(b)- n))
		if tempValue,ok := Integer(temp);ok{
			return getSquare(tempValue)
		}
	}
	return -1
}

func getSquare(n int) int{
	return n*n
}
func Integer(n float64) (int,bool){
	intN := int(n)
	return intN, n == float64(intN)
}