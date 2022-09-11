package main

import (
	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

func main() {
	mergeKListsTest()
}

// lc 144 二叉树的前序遍历
func preorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// lc 94 二叉树的中序遍历
func inorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// lc 145  二叉树的后序遍历
func postorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// lc 102  二叉树的层序遍历
// 方法2为DFS,后续再补充
// https://leetcode.cn/leetbook/read/data-structure-binary-tree/xefh1i/
// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0102.Binary-Tree-Level-Order-Traversal/
func levelOrder(root *structures.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*structures.TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		length := len(queue)
		tempRes := make([]int, 0, length)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tempRes = append(tempRes, queue[i].Val)
		}
		queue = queue[length:] //记得将队列已处理的数据清除掉
		res = append(res, tempRes)
	}
	return res
}
