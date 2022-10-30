package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//ans := permute([]int{1, 2, 3})
	//fmt.Println(ans)
	//numIslandsTest()
	//closedIslandTest()
	//maxAreaOfIslandTest()
	//countSubIslandsTest()
	//fmt.Println(generateParenthesis(3))
	//fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	slidingPuzzleTest()
}

// lc46 全排列问题[前提：元素无重复不可复选]
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	// 记录选择
	option := make([]int, 0)
	// 记录路径中元素被标记为true，避免重复使用
	visited := make([]bool, len(nums))

	// 使用闭包是不想返回res了，比较麻烦
	var backtrace func(nums []int, option []int, visited []bool)
	backtrace = func(nums []int, option []int, visited []bool) {
		// 触发结束条件
		if len(option) == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				// nums[i]已经被选择过，跳过
				continue
			}
			// 做选择
			option = append(option, nums[i])
			visited[i] = true
			// 进入下一层决策树
			backtrace(nums, option, visited)
			// 取消选择
			option = option[:len(option)-1]
			visited[i] = false
		}
	}
	backtrace(nums, option, visited)
	return ans
}

// lc47 全排列[前提：元素重复不可复选]
// https://labuladong.github.io/algo/4/31/107/
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	sort.Ints(nums)
	// 记录选择
	option := make([]int, 0)
	// 记录路径中元素被标记为true，避免重复使用
	visited := make([]bool, len(nums))

	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(nums []int, option []int, visited []bool)
	backtrack = func(nums []int, option []int, visited []bool) {
		// 触发结束条件
		if len(option) == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				// nums[i]已经被选择过，跳过
				continue
			}
			// 如何设计剪枝逻辑，把这种重复去除掉？新添加的剪枝逻辑，固定相同的元素在排列中的相对位置
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				//如果前面相邻相等元素没用过则跳过,原因:
				//标准全排列算法之所以出现重复，是因为把相同元素形成的排列序列视为不同的序列，但实际上它们应该是相同的；而如果固定相同元素形成的序列顺序，当然就避免了重复。
				continue
			}
			// 做选择
			option = append(option, nums[i])
			visited[i] = true
			// 进入下一层决策树
			backtrack(nums, option, visited)
			// 取消选择
			option = option[:len(option)-1]
			visited[i] = false
		}
	}
	backtrack(nums, option, visited)
	return ans
}

// lc78 子集问题[前提：元素无重复不可复选]
// 我们通过保证元素之间的相对顺序不变来防止出现重复的子集。
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrace func(nums []int, start int, option []int)
	backtrace = func(nums []int, start int, option []int) {
		// 前序位置，每个节点的值都是一个子集
		ans = append(ans, append([]int{}, option...))
		for i := start; i < len(nums); i++ {
			// 做选择
			option = append(option, nums[i])
			// 进入下一层决策树
			backtrace(nums, i+1, option)
			// 取消选择
			option = option[:len(option)-1]
		}
	}
	backtrace(nums, 0, option)
	return ans
}

// lc77 组合  大小为k的组合即为[1,n]回溯树中大小为k的子集
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	if n <= 0 || k <= 0 {
		return ans
	}
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(n int, start int, option []int)
	backtrack = func(n int, start int, option []int) {
		if len(option) == k {
			ans = append(ans, append([]int{}, option...))
		}
		for i := start; i <= n; i++ {
			// 做选择
			option = append(option, i)
			// 进入下一层决策树
			backtrack(n, i+1, option)
			// 取消选择
			option = option[:len(option)-1]
		}
	}
	backtrack(n, 1, option)
	return ans
}

// lc39  组合总和 [前提:无重复元素，元素可复用]
// 输入数组无重复元素，但每个元素可以被无限次使用。
// 相当于给之前的回溯树添加了一条树枝，在遍历这棵树的过程中，一个元素可以被无限次使用
// 回溯树会永远生长下去，所以我们的递归函数需要设置合适的 base case 以结束算法，即路径和大于 target 时就没必要再遍历下去
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(nums []int, start int, option []int, sum int)
	backtrack = func(nums []int, start int, option []int, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, option...))
			return
		}
		if sum > target {
			//超过目标和,则结束
			return
		}
		for i := start; i < len(nums); i++ {
			// 做选择
			option = append(option, nums[i])
			sum += nums[i]
			// 进入下一层决策树
			backtrack(nums, i, option, sum)
			// 取消选择
			sum -= nums[i]
			option = option[:len(option)-1]
		}
	}
	var sum int
	backtrack(candidates, 0, option, sum)
	return ans
}

// lc40 组合中取和为target不重复组合元素[前提：元素不重复,不复用] 排序+剪枝
func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	sort.Ints(candidates)
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(nums []int, start int, option []int, sum int)
	backtrack = func(nums []int, start int, option []int, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, option...))
			return
		}
		if sum > target {
			//超过目标和,则结束
			return
		}
		for i := start; i < len(nums); i++ {
			// 剪枝:相邻相等元素只取第一个[弱水三千,我只取一瓢]
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			// 做选择
			option = append(option, nums[i])
			sum += nums[i]
			// 进入下一层决策树
			backtrack(nums, i+1, option, sum)
			// 取消选择
			sum -= nums[i]
			option = option[:len(option)-1]
		}
	}
	var sum int
	backtrack(candidates, 0, option, sum)
	return ans
}

// lc216 找出所有相加之和为 n 的 k 个数的组合,[前提每个元素只能使用一次]
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	// k个数  n为目标和
	// 在[1,9]范围内使用k个不同的数字,和一定要大于等于k
	if n < k {
		return ans
	}
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(targetSum int, start int, option []int, sum int)
	backtrack = func(targetSum int, start int, option []int, sum int) {
		if len(option) == k && targetSum == sum {
			ans = append(ans, append([]int{}, option...))
			return
		}
		// 避免无限递归
		if len(option) > k || sum > targetSum {
			return
		}
		for i := start; i <= 9; i++ {
			// 做选择
			option = append(option, i)
			sum += i
			// 进入下一层决策树
			backtrack(targetSum, i+1, option, sum)
			// 取消选择
			option = option[:len(option)-1]
			sum -= i
		}
	}
	var sum int
	backtrack(n, 1, option, sum)
	return ans
}

// lc90 元素重复的子集问题 [前提:元素可重不可复选]
// 剪枝，如果一个节点有多条值相同的树枝相邻，则只遍历第一条，剩下的都剪掉，不要去遍历
// 体现在代码上，需要先进行排序，让相同的元素靠在一起，如果发现 nums[i] == nums[i-1]，则跳过
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) <= 0 {
		return ans
	}
	sort.Ints(nums)
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(nums []int, start int, option []int)
	backtrack = func(nums []int, start int, option []int) {
		ans = append(ans, append([]int{}, option...))
		for i := start; i < len(nums); i++ {
			// 做选择【因为含有重复元素,此时相邻重复元素只选择第一个】
			// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			option = append(option, nums[i])
			// 进入下一层决策树
			backtrack(nums, i+1, option)
			// 取消选择
			option = option[:len(option)-1]
		}
	}
	backtrack(nums, 0, option)
	return ans
}

// lc51 N皇后问题
func solveNQueens(n int) [][]string {
	var res [][]string
	// n==1时有唯一解
	if n == 1 {
		res = append(res, []string{"Q"})
		return res
	}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	// 使用闭包是不想返回res了，比较麻烦
	var backtrack func(row int)
	// 路径：board中小于row的行都已经放置了皇后
	// 选择：第row行所有列都是放置皇后的选择
	// 结束条件：row超过board的最后一行
	backtrack = func(row int) {
		// 触发结束条件
		if row == len(board) {
			temp := make([]string, n)
			for i := 0; i < len(board); i++ {
				temp[i] = strings.Join(board[i], "")
			}
			res = append(res, temp)
			return
		}
		for col := 0; col < len(board); col++ {
			//如果该行不满足皇后约束条件，则跳过本次循环
			if !isValidElement(board, row, col) {
				continue
			}
			// 满足则放置一个皇后
			board[row][col] = "Q"
			// 进入下一层决策树
			backtrack(row + 1)
			// 回溯
			board[row][col] = "."
		}
	}
	backtrack(0)
	return res
}
func isValidElement(board [][]string, row, col int) bool {
	//单层搜索时，每一层只会选择一行中的一个元素，所以不需要考虑行重复
	n := len(board)
	// 不能同列
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 不能同斜线
	// 1)左上方斜线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	// 2)右上方斜线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
func numIslandsTest() {
	grid := make([][]byte, 0)
	rank0 := []byte{'1', '1', '1', '1', '0'}
	rank1 := []byte{'1', '1', '0', '1', '0'}
	rank2 := []byte{'1', '1', '0', '0', '0'}
	rank3 := []byte{'0', '0', '0', '0', '0'}
	grid = append(grid, rank0, rank1, rank2, rank3)
	fmt.Println(numIslands(grid))
}

// lc200  岛屿数量
func numIslands(grid [][]byte) int {
	var res int
	m, n := len(grid), len(grid[0])
	// 遍历grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				//每发现1个岛屿，数量加1后将周围的岛屿全部淹没,类似visit
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// dfs 从(i,j)开始将与之相邻的陆地变为海水
func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	//将自身及周围上下左右都淹没
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func closedIslandTest() {
	grid := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0, 1, 1, 0},
		[]int{1, 0, 1, 0, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0, 1, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 0},
	}
	fmt.Println(closedIsland(grid))
}

// lc1254 封闭岛屿数量[把四周边全部淹掉]
func closedIsland(grid [][]int) int {
	var res int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		// 把左靠边、右靠边的淹掉
		dfs1254(grid, i, 0)
		dfs1254(grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		// 把上靠边、下靠边的淹掉
		dfs1254(grid, 0, j)
		dfs1254(grid, m-1, j)
	}
	// 遍历grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				//每发现1个岛屿，数量加1后将周围的岛屿全部淹没,类似visit
				dfs1254(grid, i, j)
			}
		}
	}
	return res
}

// dfs 从(i,j)开始将与之相邻的陆地变为海水
func dfs1254(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	//将自身及周围上下左右都淹没
	grid[i][j] = 1
	dfs1254(grid, i-1, j)
	dfs1254(grid, i+1, j)
	dfs1254(grid, i, j-1)
	dfs1254(grid, i, j+1)
}

func maxAreaOfIslandTest() {
	grid := [][]int{
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

// lc695 计算岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	var res int
	m, n := len(grid), len(grid[0])
	// 遍历grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 淹没岛屿，并更新最大岛屿面积
			if grid[i][j] == 1 {
				res = maxVal(dfsArea(grid, i, j), res)
			}
		}
	}
	return res
}

// dfs 从(i,j)开始将与之相邻的陆地变为海水,并返还陆地面积
func dfsArea(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	//将自身淹没
	grid[i][j] = 0
	return dfsArea(grid, i-1, j) +
		dfsArea(grid, i+1, j) +
		dfsArea(grid, i, j-1) +
		dfsArea(grid, i, j+1) + 1
}

func maxVal(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func countSubIslandsTest() {
	grid1 := [][]int{
		[]int{1, 1, 1, 0, 0},
		[]int{0, 1, 1, 1, 1},
		[]int{0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{1, 1, 0, 1, 1},
	}
	grid2 := [][]int{
		[]int{1, 1, 1, 0, 0},
		[]int{0, 0, 1, 1, 1},
		[]int{0, 1, 0, 0, 0},
		[]int{1, 0, 1, 1, 0},
		[]int{0, 1, 0, 1, 0},
	}
	fmt.Println(countSubIslands(grid1, grid2))
}

// lc1905 统计子岛屿 反过来说，如果岛屿 B 中存在一片陆地，在岛屿 A 的对应位置是海水，那么岛屿 B 就不是岛屿 A 的子岛。
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var res int
	m, n := len(grid1), len(grid1[0])
	// 将岛屿2在岛屿1中不是岛屿的全部淹掉
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs1905(grid2, i, j)
			}
		}
	}
	// 重新遍历grid2,计算岛屿的个数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				//每发现1个岛屿，数量加1后将周围的岛屿全部淹没,类似visit
				dfs1905(grid2, i, j)
			}
		}
	}
	return res
}

// dfs 从(i,j)开始将与之相邻的陆地变为海水
func dfs1905(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	//将自身及周围上下左右都淹没
	grid[i][j] = 0
	dfs1905(grid, i-1, j)
	dfs1905(grid, i+1, j)
	dfs1905(grid, i, j-1)
	dfs1905(grid, i, j+1)
}

// lc22 生成括号
//一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解。
//对于一个「合法」的括号字符串组合 p，必然对于任何  0 <= i < len(p) 都有：子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n <= 0 {
		return ans
	}
	// 做选择
	option := make([]string, 0)
	//使用闭包是不想返回ans,比较麻烦
	var backtrack func(left int, right int, option []string)
	backtrack = func(left int, right int, option []string) {
		// 性质2
		if right < left {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			ans = append(ans, strings.Join(option, ""))
			//fmt.Println(strings.Join(option, ""))
			return
		}
		//尝试放(
		option = append(option, "(")
		backtrack(left-1, right, option)
		option = option[:len(option)-1]

		//尝试放)
		option = append(option, ")")
		backtrack(left, right-1, option)
		option = option[:len(option)-1]
	}
	backtrack(n, n, option)
	return ans
}

// lc752 BFS 最小转动次数
// TODO: 优化空间双向BFS
func openLock(deadends []string, target string) int {
	visit := map[string]bool{} //保存被访问的数字,避免回头循环
	deadMap := make(map[string]bool, len(deadends))
	for _, dead := range deadends {
		deadMap[dead] = true
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visit["0000"] = true
	step := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cutStr := queue[i]
			// 判断是否到达终点及黑名单
			if deadMap[cutStr] {
				continue
			}
			if cutStr == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := plusOne(cutStr, j)
				if !visit[up] {
					visit[up] = true
					queue = append(queue, up)
				}
				down := minusOne(cutStr, j)
				if !visit[down] {
					visit[down] = true
					queue = append(queue, down)
				}
			}
		}
		queue = queue[n:]
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

func plusOne(s string, i int) string {
	strByte := []byte(s)
	if strByte[i] == '9' {
		strByte[i] = '0'
	} else {
		strByte[i] += 1
	}
	return string(strByte)
}

func minusOne(s string, i int) string {
	strByte := []byte(s)
	if strByte[i] == '0' {
		strByte[i] = '9'
	} else {
		strByte[i] -= 1
	}
	return string(strByte)
}

func slidingPuzzleTest() {
	board := [][]int{
		[]int{4, 1, 2},
		[]int{5, 0, 3},
	}
	fmt.Println(slidingPuzzle(board))
}

// lc773 数字华容道
//TODO: 计算初始面板相邻坐标索引可以使用函数计算,不需要固定法
func slidingPuzzle(board [][]int) int {
	const target = "123450"
	boardByte := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			boardByte = append(boardByte, '0'+byte(v))
		}
	}
	start := string(boardByte)
	if start == target {
		return 0
	}
	step := 0
	visit := map[string]bool{}
	queue := make([]string, 0)
	queue = append(queue, start)
	visit[start] = true
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curNum := queue[i]
			// 检查是否到达终点
			if curNum == target {
				return step
			}
			// 0的相邻点,其实就是棋局里原有位置的相邻节点
			zeroIndex, zeroNeighbours := getZeroNeighbours(curNum)
			for _, adjIndex := range zeroNeighbours {
				newNum := changeNum(curNum, zeroIndex, adjIndex)
				if !visit[newNum] {
					visit[newNum] = true
					queue = append(queue, newNum)
				}
			}

		}
		queue = queue[n:]
		step++
	}
	return -1
}

func changeNum(str string, i int, j int) string {
	strByte := []byte(str)
	temp := strByte[i]
	strByte[i] = strByte[j]
	strByte[j] = temp
	return string(strByte)
}

// 返回0的索引及对应原有棋盘该节点的所有相邻点
func getZeroNeighbours(num string) (int, []int) {
	var zeroIndex int
	numByte := []byte(num)
	for i := 0; i < len(numByte); i++ {
		if numByte[i] == '0' {
			zeroIndex = i
			break
		}
	}
	neighbours := [6][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	return zeroIndex, neighbours[zeroIndex]
}

// lc32 解数独回溯
func solveSudoku(board [][]byte) {
	backtrackBoard(board, 0, 0)
}
func backtrackBoard(board [][]byte, i int, j int) bool {
	m, n := 9, 9
	if j == n {
		//穷举到最后一列则进行下一行
		return backtrackBoard(board, i+1, 0)
	}
	if i == m {
		// 找到可行解，触发base case
		return true
	}
	if board[i][j] != '.' {
		return backtrackBoard(board, i, j+1)
	}
	possibleChar := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, ch := range possibleChar {
		if !isValidBoard(board, i, j, ch) {
			continue
		}
		board[i][j] = ch
		//如果找到一个可行解,立即结束
		if backtrackBoard(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	//穷举完仍没有则此路不通
	return false
}

// 判断 board[r][c] 是否可以填入ch
func isValidBoard(board [][]byte, r int, c int, ch byte) bool {

	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[r][i] == ch {
			return false
		}
		// 判断列是否存在重复
		if board[i][c] == ch {
			return false
		}
		// 判断 3 x 3 方框是否存在重复
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == ch {
			return false
		}
	}
	return true
}
