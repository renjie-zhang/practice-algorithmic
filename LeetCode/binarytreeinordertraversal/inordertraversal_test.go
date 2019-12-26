package binarytreeinordertraversal

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tree1 := &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 2}
	tree3 := &TreeNode{Val: 3}
	tree1.Right = tree2
	tree2.Left = tree3
	//temp := inorderTraversal(tree1)
	temp := inorderTraversal2(tree2)
	fmt.Println(temp)
}
