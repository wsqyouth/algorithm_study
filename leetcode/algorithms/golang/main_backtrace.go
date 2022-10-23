package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	ans := permute([]int{1, 2, 3})
	fmt.Println(ans)
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
