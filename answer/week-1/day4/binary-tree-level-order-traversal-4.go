package day4

import "container/list"

//方法一BFS广度优先搜索
//方法二深度优先搜索
func levelOrderBranch(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue list.List
	queue.PushFront(root)
	for queue.Len() > 0 {
		var levelRes []int
		n := queue.Len()
		for i := 0; i < n; i++ {
			var treeNode = queue.Remove(queue.Front())
			levelRes = append(levelRes, treeNode.(*TreeNode).Val)
			if treeNode.(*TreeNode).Left != nil {
				queue.PushBack(treeNode.(*TreeNode).Left)
			}
			if treeNode.(*TreeNode).Right != nil {
				queue.PushBack(treeNode.(*TreeNode).Right)
			}
		}
		res = append(res, levelRes)
	}
	return res
}
func levelDep(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var resMap = make(map[int][]int)
	var res [][]int
	helper(root, 0, resMap)
	for i := 0; i < len(resMap); i++ {
		res = append(res, resMap[i])
	}
	return res
}
func helper(root *TreeNode, level int, resMap map[int][]int) {
	if root != nil {
		resMap[level] = append(resMap[level], root.Val)
		helper(root.Left, level+1, resMap)
		helper(root.Right, level+1, resMap)
	}
}
