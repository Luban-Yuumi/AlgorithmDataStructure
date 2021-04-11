package day3




//时间复杂度为o(n)因为每个节点都会访问一遍
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	right := lowestCommonAncestor(root.Right, p, q)
	left := lowestCommonAncestor(root.Left, p, q)

	if right == nil {
		return left
	} else if left == nil {
		return right
	} else {
		return root
	}
}
