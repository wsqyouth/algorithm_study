package main

import (
	"fmt"
	"math"
)

func main() {
	//climbStairsTest()
	fmt.Println(coinChange([]int{3, 2, 5}, 11))
	//fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	// fmt.Println(isPalindrome(123))
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

// lc746 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	//定义：到达第i个台阶所花费的最少体力为dp[i]
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	// 最后一步不需要花费，因此取倒数第1,2步最小值
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

// lc322 零钱兑换 使用迭代
func coinChange(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// dp数组初始化为特殊值
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 外层遍历所有状态的取值
	for i := 0; i < len(dp); i++ {
		// 内层计算所有选择的最小值
		for _, coin := range coins {
			// fmt.Println(coin)
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	// 检查是否存在
	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}
	return -1
}

// lc322 零钱兑换 完全背包
func coinChangeOld(coins []int, amount int) int {
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

// lc300 最长递增子序列
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

// lc53 最大的子数组和
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	//dp:以 nums[i] 为结尾的「最大子数组和」为 dp[i]。
	dp := make([]int, length)
	// base case
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i]) //要么自身自成一派,要么和前面元素连接
	}
	var res = math.MinInt32
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}

// 时间复杂度O(logn),空间复杂度O(1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

// lc6 判断回文数 基本原理是掌握数字字符串转整型的思路
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temp := x
	var reverseRes int
	for temp > 0 {
		lastNum := temp % 10
		temp = temp / 10
		reverseRes = reverseRes*10 + lastNum
	}
	return reverseRes == x
}
