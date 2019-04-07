package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一直接中序遍历获得遍历后的结果
// 然后将排序后的结果
// 与 返回的结果做比较
// 这种做法 由于要存储所有遍历结果 空间复杂度比较差
//方法二：每次保存上一个节点的值 与当前值做比较 看是不是升序
func main() {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	var prev = []*TreeNode{nil}
	return helper(root, prev)
}

func helper(root *TreeNode, prev []*TreeNode) bool {
	if root == nil {
		return true
	}

	if !helper(root.Left, prev) {
		return false
	}
	if prev[0] != nil && root.Val <= prev[0].Val {
		return false
	}
	//prev = root 这样赋值不会改变函数外prev
	prev[0] = root
	return helper(root.Right, prev)
}
