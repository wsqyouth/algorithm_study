package main

import (
	"fmt"

	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

type TreeNode = structures.TreeNode

func main() {
	maxDepthTest()
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

// lc104  二叉树的最大深度
// 思路：这道题是树思想的体现,能够讲清楚递归的本质
// 1. 自顶向下 -> 先操作根节点（状态已知）-> 前序遍历
// 2. 自底向上 -> 后操作根节点（状态可能未知）-> 后续遍历
// 前序位置的代码执行是自顶向下的，而后序位置的代码执行是自底向上的
// 意味着前序位置的代码只能从函数参数中获取父节点传递来的数据，而后序位置的代码不仅可以获取参数数据，还可以获取到子树通过函数返回值传递回来的数据
// https://leetcode.cn/leetbook/read/data-structure-binary-tree/xefb4e/
// https://labuladong.github.io/algo/2/21/36/
func maxDepth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}

	// 自底向上,bottomToTop,计算每个节点以下的情况，后续遍历
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func maxDepthTest() {
	preOrder := []int{3, 9, structures.NULL, structures.NULL, 20, 15, 7} //前序:根左右
	inOrder := []int{structures.NULL, 9, structures.NULL, 3, 15, 20, 7}  //中序:左根右
	root := structures.PreIn2Tree(preOrder, inOrder)                     //前序后续构建树
	structures.PrintTree(root)
	fmt.Println(maxDepth(root))

}

// lc543  二叉树的直径
// 思路：解决这题的关键在于，每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和。
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	checkDiameter(root, &res)
	return res
}

func checkDiameter(root *TreeNode, result *int) int {
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if root == nil {
		return 0
	}
	leftDepth := checkDiameter(root.Left, result)
	rightDepth := checkDiameter(root.Right, result)
	diameter := leftDepth + rightDepth
	*result = maxFunc(diameter, *result)
	return maxFunc(leftDepth, rightDepth) + 1
}

// lc 112 判断树中是否存在根节点到叶子节点路径和为target
// 参考：https://leetcode.cn/leetbook/read/data-structure-binary-tree/xo566j/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// lc 102 判断二叉树是否轴对称
// 参考：https://leetcode.cn/leetbook/read/data-structure-binary-tree/xoxzgv/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}
func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	// 二者任意一个为空，另一个不为空的场景,妙
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	// 树是对称的,最左侧和最右侧比较
	//左子节点的左子节点和右子节点的右子节点比较，左子节点的右子节点和右子节点的左子节点比较
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
