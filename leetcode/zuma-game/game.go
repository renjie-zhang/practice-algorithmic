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
package zuma_game

import "fmt"


func FindMinStep(board string, hand string) int {
	minStep := len(hand)+1
	hands := make(map[byte]int)
	for i:=0; i<len(hand); i++ {
		hands[hand[i]] ++
	}
	dfs(board, hand, hands, 0, &minStep)

	if minStep==len(hand)+1 {
		return -1
	}
	return minStep
}

func dfs(board string, hand string, hands map[byte]int, step int, pMinStep *int) {
	if step==len(hand) {
		return
	}
	if step>=*pMinStep {
		return
	}

	for i, j:=0, 0; j<len(board); i, j=j+1, j+1 {
		for j<len(board)-1&&board[j+1]==board[j] {
			j++
		}

		// TODO: Code improvement for if-else combination

		dupN := j-i+1 // Duplicated balls number
		rh := hands[board[i]] // Remaining hand of the color
		if dupN==1 {
			if rh>=2 {
				hands[board[i]] -= 2
				board2 := insertBoard(board, i, board[i], 2)
				board2 = eliminate(board2, i)
				if len(board2)==0 {
					if step+2<*pMinStep {
						*pMinStep = step+2
					}
					break
				}
				dfs(board2, hand, hands, step+2, pMinStep)
				hands[board[i]] += 2
			}
		} else { // dupN>=2
			if rh>=1 {
				hands[board[i]] -= 1
				board2 := insertBoard(board, i, board[i], 1)
				board2 = eliminate(board2, i)
				if len(board2)==0 {
					if step+1<*pMinStep {
						*pMinStep = step+1
					}
					break
				}
				dfs(board2, hand, hands, step+1, pMinStep)
				hands[board[i]] += 1
			} else { // rh==0
				// Insert different colors
				for c, num := range hands {
					if c==board[i] || num==0 {
						continue
					}

					hands[c] -= 1
					board2 := insertBoard(board, i+1, c, 1)
					dfs(board2, hand, hands, step+1, pMinStep)
					hands[c] += 1
				}
			}
		} // endof if-else
	} // endof for
}

func insertBoard(board string, i int, c byte, n int) string {
	newBoard := board[0:i]
	for ;n>0; n-- {
		newBoard += string(c)
	}
	newBoard += board[i:]
	return newBoard
}

func eliminate(board string, i int) string {
	for i>=0&&i<=len(board)-1{
		l, r := i, i
		for (l>0)&&(board[l-1]==board[i]) {
			l--
		}
		for (r<len(board)-1)&&(board[r+1]==board[i]) {
			r++
		}
		if r-l<2 {
			break
		}
		board = fmt.Sprintf("%s%s", board[0:l], board[r+1:])
		i = l
	}
	return board
}
