package day5

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return 1 + minDepth(root.Left) + minDepth(root.Right)
	}

	return 1 + int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))))

}
