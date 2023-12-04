type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输出二叉树的左视图
func leftView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		result = append(result, queue[0].Val) // 添加每一层的第一个元素
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

// 输出二叉树的右视图
func rightView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		result = append(result, queue[size-1].Val) // 添加每一层的最后一个元素
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}