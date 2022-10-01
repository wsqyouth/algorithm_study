package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

type TreeNode = structures.TreeNode

//func main() {
//sortArrayTest()
//findKthLargestTest()
//topKFrequentTest()
//}

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

// lc107 层次遍历从底向上打印
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
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
	// 也可以在每一个切片插入时从头部插入，go没有这样的特性只能手动数组翻转下
	length := len(res)
	for i := 0; i < (length)/2; i++ {
		res[i], res[length-i-1] = res[length-i-1], res[i]
	}
	return res
}

// lc199 二叉树的右视图 判断是否遍历到单层的最后面的元素，如果是，就放进result数组中，随后返回result就可以了
//func rightSideView(root *TreeNode) []int {}
// lc637 二叉树每层求平均值
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	res := make([]float64, 0)
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
		var sum float64
		for _, val := range tempRes {
			sum += float64(val)
		}
		tempAverageVal := sum / float64(length)
		res = append(res, tempAverageVal)
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
	root := structures.PreIn2Tree(preOrder, inOrder)                     //前序中序构建树
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

func buildTreeTest() {
	preOrder := []int{3, 9, 20, 15, 7}   //前序:根左右
	inOrder := []int{9, 3, 15, 20, 7}    //中序:左根右
	root := buildTree(preOrder, inOrder) //前序中序构建树
	structures.PrintTree(root)
}

// lc 105 前序中序构建二叉树
// 参考：https://labuladong.github.io/algo/2/21/38/
func buildTree(preorder []int, inorder []int) *TreeNode {
	//用preorder和inorder构造二叉树
	return buildTreeHelper(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}
func buildTreeHelper(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart] //root节点
	var index int
	for i, v := range inorder {
		if v == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftSize := index - inStart
	//递归构造左右子树
	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1)
	root.Right = buildTreeHelper(preorder,
		preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)
	return root
}

func buildTreeTest106() {
	inOrder := []int{9, 3, 15, 20, 7}        //中序:左根右
	postOrder := []int{9, 15, 7, 20, 3}      //后序:左右根
	root := buildTree106(inOrder, postOrder) //中序后序构建树
	structures.PrintTree(root)
}

// lc 106 后序中序构建二叉树
// 参考：https://labuladong.github.io/algo/2/21/38/
func buildTree106(inorder []int, postorder []int) *TreeNode {
	//用postorder和inorder构造二叉树
	return buildTreeHelper106(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1)
}
func buildTreeHelper106(inorder []int, inStart int, inEnd int,
	postorder []int, postStart int, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[postEnd] //root节点
	var index int
	for i, v := range inorder {
		if v == rootVal {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftSize := index - inStart
	//递归构造左右子树
	root.Left = buildTreeHelper106(inorder, inStart, index-1,
		postorder, postStart, postStart+leftSize-1)
	root.Right = buildTreeHelper106(inorder, index+1, inEnd,
		postorder, postStart+leftSize, postEnd-1)
	return root
}

// lc 297题 二叉树序列化和反序列化
// 思想: https://labuladong.github.io/algo/2/21/39/
// 参考:https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0297.Serialize-and-Deserialize-Binary-Tree/
func SerializesTest() {
	inOrder := []int{9, 3, 15, 20, 7}        //中序:左根右
	postOrder := []int{9, 15, 7, 20, 3}      //后序:左右根
	root := buildTree106(inOrder, postOrder) //中序后序构建树

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	fmt.Println(data)
	structures.PrintTree(ans)
}

type Codec struct {
	builder strings.Builder
	input   []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.builder.WriteString("#,")
		return ""
	}
	this.builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.input = strings.Split(data, ",")
	return this.deserializeHelper()
}

// 想象为满二叉树进行前序遍历处理
func (this *Codec) deserializeHelper() *TreeNode {
	if this.input[0] == "#" {
		this.input = this.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.deserializeHelper(),
		Right: this.deserializeHelper(),
	}
}

func constructMaximumBinaryTreeTest() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	structures.PrintTree(root)

}

// lc654 构建最大二叉树,其实就是按照切片的顺序前序遍历构造即可，和lc297几乎一样
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxNumIndex, maxNum := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxNumIndex = i
		}
	}
	return &TreeNode{
		Val:   maxNum,
		Left:  constructMaximumBinaryTree(nums[:maxNumIndex]),
		Right: constructMaximumBinaryTree(nums[maxNumIndex+1:]),
	}
}

func sortArrayTest() {
	nums := []int{7, 8, 4, 3, 2, 1, 6, 0, 5}
	ans := sortArray(nums)
	fmt.Println(ans)
	//quickSort
	quickSortArray(nums)
	fmt.Println((nums))
	/*
		fmt.Println(mergeSortArray([]int{1, 2, 3}, []int{4, 5, 6}))
	*/
}

// lc912 升序排序,使用归并排序,二叉树后续遍历
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	return sortArrayHelper(nums, 0, len(nums)-1)
}
func sortArrayHelper(nums []int, low int, high int) []int {
	if low == high {
		return []int{nums[low]}
	}
	mid := low + (high-low)/2
	left := sortArrayHelper(nums, low, mid)
	right := sortArrayHelper(nums, mid+1, high)

	return mergeSortArray(left, right)
}
func mergeSortArray(left []int, right []int) []int {
	m, n, length := len(left), len(right), len(left)+len(right)
	res := make([]int, length)
	var i, j, k int
	for i < m && j < n {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	res = res[0:k]
	if i < m {
		res = append(res, left[i:m]...)
	}
	if j < n {
		res = append(res, right[j:n]...)
	}
	return res
}

func quickSortArray(nums []int) {
	if len(nums) == 0 {
		return
	}
	sortQuickArrayHelper(nums, 0, len(nums)-1)
}

// 快排排序，二叉树前序遍历
func sortQuickArrayHelper(nums []int, low int, high int) {
	if low >= high {
		return
	}
	// 对nums进行切分,使得nums[low..p-1] <= nums[p] < nums[p+1,high]
	p := partitionArray(nums, low, high)
	sortQuickArrayHelper(nums, low, p-1)
	sortQuickArrayHelper(nums, p+1, high)
}

//ref: https://mp.weixin.qq.com/s/8ZTMhvHJK_He48PpSt_AmQ
func partitionArray(nums []int, low int, high int) int {
	// low对应基准。 i,j为开区间，保证覆盖到[low,high]
	pivot := nums[low]
	i, j := low+1, high
	for i <= j {
		for i < high && nums[i] <= pivot {
			i++
		}
		for j > low && nums[j] > pivot {
			j--
		}
		// 以上步骤使得[low,i) <= pivlot && (j,high] > pivot
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, low, j)
	return j
}
func swap(nums []int, i int, j int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}

// lc215 数组第k大元素
func findKthLargestTest() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}

func findKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1
	k = len(nums) - k //转换为排名第k'的元素

	for low <= high {
		// 对nums进行切分,使得nums[low..p-1] <= nums[p] < nums[p+1,high]
		p := partitionArray(nums, low, high)
		if p < k {
			low = p + 1
		} else if p > k {
			high = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

// 还缺一个堆排序,能够对比这3个算法的不同
// lc347 前k个高频元素,可以使用小顶堆实现，这里仍然使用快排partition算法
func topKFrequentTest() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}

type NumFreq struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	// 1、构建数字出现频率的列表
	lenNum := len(nums)
	if lenNum == 1 {
		return nums
	}
	mapNum := make(map[int]int)
	for i := 0; i < lenNum; i++ {
		mapNum[nums[i]]++
	}
	numFreqList := []NumFreq{}
	for i, v := range mapNum {
		numFreqList = append(numFreqList, NumFreq{i, v})
	}
	// 2.使用partition算法进行查找
	var res []int
	low, high := 0, len(numFreqList)-1
	for low <= high {
		// 对nums进行切分,使得nums[low..p-1] <= nums[p] < nums[p+1,high]
		p := partitionStructArray(numFreqList, low, high)
		k := p - low + 1 //表示基准左侧元素的个数
		if p < k {
			low = p + 1
		} else if p > k {
			high = p - 1
		} else {
			//找到结果后填充
			for ii := low; ii <= p; ii++ {
				res = append(res, numFreqList[ii].num)
			}
		}
	}
	return res
}
func partitionStructArray(nums []NumFreq, low int, high int) int {
	// low对应基准。 i,j为开区间，保证覆盖到[low,high]
	pivot := nums[low].num
	i, j := low+1, high
	for i <= j {
		for i < high && nums[i].num <= pivot {
			i++
		}
		for j > low && nums[j].num > pivot {
			j--
		}
		// 以上步骤使得[low,i) <= pivlot && (j,high] > pivot
		if i >= j {
			break
		}
		swapNum(nums, i, j)
	}
	swapNum(nums, low, j)
	return j
}

func swapNum(nums []NumFreq, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
