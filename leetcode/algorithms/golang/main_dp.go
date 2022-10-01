package main

import (
	"fmt"
	"math"
)

func main() {
	//climbStairsTest()
	//fmt.Println(coinChange([]int{3, 2, 5}, 11))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func climbStairsTest() {
	fmt.Println(climbStairs(3))
}

// lc70 爬楼梯
func climbStairs(n int) int {
	// 备忘录
	memo := make([]int, n+1)
	return dp(n, memo)
}

// 定义：爬到第n阶台阶的方法的各位为dp(n)
func dp(n int, memo []int) int {
	//base case
	if n <= 2 {
		return n
	}
	if memo[n] > 0 {
		return memo[n]
	}
	//状态转移方程：爬到第n级等于爬到第n-1级方法与爬到第n-2层的方法个数之和
	memo[n] = dp(n-1, memo) + dp(n-2, memo)
	return memo[n]
}

// lc322 零钱兑换 完全背包
func coinChange(coins []int, amount int) int {
	// 备忘录
	memo := make([]int, amount+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = math.MinInt32 //特殊值
	}
	return dpCoin(coins, amount, memo)
}
func dpCoin(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != math.MinInt32 {
		return memo[amount] //避免重复计算
	}
	res := math.MaxInt32
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dpCoin(coins, amount-coin, memo)
		if subProblem == -1 {
			continue
		}
		//在子问题中选择最优解，然后加1
		res = min(res, subProblem+1)
	}
	// 计算结果存入备忘录
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func lengthOfLIS(nums []int) int {
	// dp[i]: 以nums[i]结尾最长递增子序列【不连续】的长度
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	var res int
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}
