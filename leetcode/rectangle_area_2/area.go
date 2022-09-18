package rectangle_area_2

import "sort"

func rectangleArea(rectangles [][]int) (ans int) {
	n := len(rectangles) * 2
	hBound := make([]int, 0, n)
	for _, r := range rectangles {
		hBound = append(hBound, r[1], r[3])
	}
	// 排序，方便下面去重
	sort.Ints(hBound)
	m := 0
	for _, b := range hBound[1:] {
		if hBound[m] != b {
			m++
			hBound[m] = b
		}
	}
	hBound = hBound[:m+1]

	type tuple struct{ x, i, d int }
	sweep := make([]tuple, 0, n)
	for i, r := range rectangles {
		sweep = append(sweep, tuple{r[0], i, 1}, tuple{r[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool { return sweep[i].x < sweep[j].x })

	seg := make([]int, m)
	for i := 0; i < n; i++ {
		j := i
		for j+1 < n && sweep[j+1].x == sweep[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		// 一次性地处理掉一批横坐标相同的左右边界
		for k := i; k <= j; k++ {
			idx, diff := sweep[k].i, sweep[k].d
			left, right := rectangles[idx][1], rectangles[idx][3]
			for x := 0; x < m; x++ {
				if left <= hBound[x] && hBound[x+1] <= right {
					seg[x] += diff
				}
			}
		}
		cover := 0
		for k := 0; k < m; k++ {
			if seg[k] > 0 {
				cover += hBound[k+1] - hBound[k]
			}
		}
		ans += cover * (sweep[j+1].x - sweep[j].x)
		i = j
	}
	return ans % (1e9 + 7)
}
