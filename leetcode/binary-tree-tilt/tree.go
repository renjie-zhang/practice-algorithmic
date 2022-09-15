package binary_tree_tilt

func findTilt(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	var answer int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		answer += abs(left - right)
		return left + right + node.Val

	}
	dfs(root)
	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}
