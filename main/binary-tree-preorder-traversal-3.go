package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(preorderTraversal(root))
}
func preorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}
	res = append(res, root.Val)
	return append(append(res, preorderTraversal(root.Left)...), preorderTraversal(root.Right)...)
	//return append(append(preorderTraversal(root.Left), res...), preorderTraversal(root.Right)...)中序
	//return append(append(preorderTraversal(root.Left),preorderTraversal(root.Right)...),res...)后序
}
