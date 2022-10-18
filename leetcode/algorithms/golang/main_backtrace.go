package main

import (
	"fmt"
)

func main() {

	ans := permute([]int{1, 2, 3})
	fmt.Println(ans)
}

// lc46 全排列问题
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

// lc78 子集组合问题
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	// 记录选择
	option := make([]int, 0)
	// 使用闭包是不想返回res了，比较麻烦
	var backtrace func(nums []int, start int, option []int)
	backtrace = func(nums []int, start int, option []int) {
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
