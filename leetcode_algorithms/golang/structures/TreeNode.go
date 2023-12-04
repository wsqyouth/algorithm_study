package structures

import (
	"fmt"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 层次遍历结果构建二叉树
func Ints2TreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	//构建队列存树节点,为后面一层做准备
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:] //每次拆第一个作为子树根

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func indexOf(nums []int, val int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	msg := fmt.Sprintf("%d not exist in %v", val, nums)
	panic(msg)
}

// PreIn2Tree 将preOrder和inOrder转换为二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("pre in are not same length")
	}
	if len(in) == 0 || len(pre) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}

	if len(in) == 1 {
		return root
	}

	idx := indexOf(in, root.Val)

	root.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	root.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return root
}

// Tree2PreOrder 将二叉树转化为preOrder切片
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2PreOrder(root.Left)...)
	res = append(res, Tree2PreOrder(root.Right)...)
	return res
}

// Tree2InOrder 将二叉树转化为inOrder切片
func Tree2InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := Tree2InOrder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2InOrder(root.Right)...)
	return res
}

// Tree2PostOrder 将二叉树转化为postOrder切片
func Tree2PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2PostOrder(root.Left)
	res = append(res, Tree2PostOrder(root.Right)...)
	res = append(res, root.Val)
	return res
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		length := len(queue)
		//处理这一层级的二叉树，将该层级节点的所有节点保存在queue中
		for _, v := range queue {
			//fmt.Printf("  %v", v.Val)
			if v.Val == NULL {
				fmt.Printf("  %v", "null")
			} else {
				fmt.Printf("  %v", v.Val)
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		fmt.Printf("\n")
		//删掉已处理的层级的节点值
		queue = queue[length:]
	}
}
