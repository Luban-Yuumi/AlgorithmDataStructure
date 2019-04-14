package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return 1 + minDepth(root.Left) + minDepth(root.Right)
	} else {
		return 1 + int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))))
	}
}
