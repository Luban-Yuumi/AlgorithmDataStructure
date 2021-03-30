package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	treeHelper(&res, root)
	return res
}

func treeHelper(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	treeHelper(res, root.Left)
	treeHelper(res, root.Right)
}

