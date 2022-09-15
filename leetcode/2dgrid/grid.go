package main

import "fmt"

func main() {
	var t = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	grid := shiftGrid(t, 1)
	fmt.Println(grid)
}

func shiftGrid(grid [][]int, k int) [][]int {
	var col = len(grid[0])
	var row = len(grid)
	var result = make([][]int, row)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, col)
	}
	// 如果 k > m*n
	var count = col * row
	if k == count {
		// 移动循环回到了初始点
		return grid
	} else if k < count {
		k = k % count
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			inde := (i*col + j + k) % (col * row)
			result[inde/col][inde%col] = grid[i][j]
		}
	}
	fmt.Println(result)
	return result
}

func location(x, y, k, col, row int) (int, int) {
	var resultX int
	var resultY int
	var temp = x + k + 1
	if temp > col {
		var t = temp % col
		var i = temp / col
		resultY = y + i

		if resultY > row {
			resultY = resultY % row
		}
		resultX = x + t
	} else if temp <= col {
		resultY = y
		resultX = x + k
	}
	return resultX, resultY
}
