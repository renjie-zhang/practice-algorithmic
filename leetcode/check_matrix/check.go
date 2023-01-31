package check

func CheckXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == j {
				if grid[i][j] != 0 {
					continue
				} else {
					return false
				}
			}
			if i+j == len(grid[i])-1 {
				if grid[i][j] != 0 {
					continue
				} else {
					return false
				}

			}
			if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
