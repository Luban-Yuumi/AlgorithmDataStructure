package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, make([]int, 0, 1))
}

func helper(root *TreeNode, prev []int) bool {
	if root == nil {
		return true
	}

	if !helper(root.Left, prev) {
		return false
	}
	fmt.Println(prev)

	if len(prev) != 0 && prev[0] >= root.Val {
		return false
	}

	if len(prev) == 0 {
		prev = append(prev, root.Val)
	} else {
		prev[0] = root.Val
	}
	fmt.Println(root)
	return helper(root.Right, prev)
}

func main() {
	left := &TreeNode{
		Val: 1,
	}
	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: nil,
	}

	fmt.Println(isValidBST(root))
}
