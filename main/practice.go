package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre = make([]*TreeNode, 1)
	return helper(root, pre)
}

func helper(root *TreeNode, prev []*TreeNode) bool {
	if root == nil {
		return true
	}

	if !helper(root.Left, prev) {
		return false
	}

	if prev != nil && prev[0].Val >= root.Val {
		return false
	}

	prev[0] = root

	return helper(root.Right, prev)
}
