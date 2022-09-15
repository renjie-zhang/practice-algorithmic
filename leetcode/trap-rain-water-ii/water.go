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
package trap_rain_water_ii

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])
	var resutl int

	if m < 2 || n < 2{
		return 0
	}
	// 记录已经访问的点
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	// 将外围一圈标为已访问
	h := &hp{}
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{v, i, j})
				vis[i][j] = true
			}
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && !vis[nx][ny] {
				if heightMap[nx][ny] < cur.h {
					resutl += cur.h - heightMap[nx][ny]
				}
				vis[nx][ny] = true
				heap.Push(h, cell{max(heightMap[nx][ny], cur.h), nx, ny})
			}
		}
	}
	return resutl
}

type cell struct {
	h,x,y int
}
type hp []cell
func (h hp) Len() int{
	return len(h)
}

func (h hp) Less(i,j int) bool{
	return h[i].h<h[j].h
}

func (h hp) Swap(i,j int){
	h[i],h[j] = h[j],h[i]
}

func (h *hp) Pop() interface{}{
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h *hp) Push(v interface{}){
	*h = append(*h,v.(cell))
}

func max(a,b int) int{
	if b>a {
		return b
	}
	return a
}




