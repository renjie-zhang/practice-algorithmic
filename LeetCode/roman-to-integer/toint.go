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
package roman_to_integer

import (
	"fmt"
	"strings"
)

// my solution
func RomanToInt(s string) int {
	values := strings.Split(s,"")
	var result = 0
	var length = len(values)
	for i := 0;i<length ;  {
		switch values[i] {
		case "I":
			if i+1 == length{
				result = result + 1
				i++
			}else {
				temp := i
				temp ++
				if values[temp] == "V"{
					result = result + 4
					i = i+2
				}else if values[temp] == "X" {
					result = result + 9
					i = i+2
				}else {
					result = result + 1
					i++
				}
			}

		case "V":
			result = result + 5
			i++
		case "X":
			if i+1 ==length{
				result = result + 10
				i++
			}else {
				temp := i
				temp ++
				if values[temp] == "L"{
					result = result + 40
					i = i+2
				}else if values[temp] == "C" {
					result = result + 90
					i = i+2
				}else {
					result = result + 10
					i++
				}
			}
		case "L":
			result = result + 50
			i++
		case "C":
			if i + 1 == length{
				result = result + 100
				i++
			}else {
				temp := i
				temp ++
				if values[temp] == "D"{
					result = result + 400
					i = i+2
				}else if values[temp] == "M" {
					result = result + 900
					i = i+2
				}else {
					result = result + 100
					i++
				}
			}
		case "D":
			result = result + 500
			i++
		case "M":
			result = result + 1000
			i++
		default:
			fmt.Printf("err")
		} 
	}
	return result
}

// other solution
func RomanToInt2(s string) int{
	values := strings.Split(s,"")
	result := 0
	preValue := getValue(values[0])
	length := len(values)
	for i := 1;i < length ;i++  {
		temp := getValue(values[i])
		if preValue < temp{
			result -= preValue
		}else{
			result += preValue
		}
		preValue = temp
	}
	result += preValue
	return result

}

func getValue(s string) int{
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}