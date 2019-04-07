package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//时间复杂度为o(n)因为每个节点都会访问一遍
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else {
		if right == nil {
			return left
		} else {
			return root
		}
	}
}
