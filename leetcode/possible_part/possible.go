package possible_part

//https://leetcode.cn/problems/possible-bipartition/

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unionFind{parent, make([]int, n)}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) union(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func (uf unionFind) isConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func PossibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	uf := newUnionFind(n)
	for x, nodes := range g {
		for _, y := range nodes {
			uf.union(nodes[0], y)
			if uf.isConnected(x, y) {
				return false
			}
		}
	}
	return true
}
