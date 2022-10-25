package shortes_bridge

// https://leetcode.cn/problems/shortest-bridge/

func shortestBridge(grid [][]int) (step int) {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			q := []pair{}
			var dfs func(int, int)
			dfs = func(i, j int) {
				grid[i][j] = -1
				q = append(q, pair{i, j})
				for _, d := range dirs {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						dfs(x, y)
					}
				}
			}
			dfs(i, j)

			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return
}
