package evaluate_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func EvaluateTree(root *TreeNode) bool {
	// 后序遍历
	if root.Left == nil {
		return root.Val == 1
	}
	if root.Val == 2 {
		return EvaluateTree(root.Left) || EvaluateTree(root.Right)
	}
	return EvaluateTree(root.Left) && EvaluateTree(root.Right)
}
