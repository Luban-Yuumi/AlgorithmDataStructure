package day3


//时间复杂度 o(n)
func lowestCommonAncestorBinaryTree(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		root = lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		root = lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
