package main

import (
	"fmt"
	"math"
)

/*
 */
func main() {
	//climbStairsTest()
	// fmt.Println(coinChange([]int{3, 2, 5}, 11))
	//fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	// fmt.Println(isPalindrome(123))
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(maxProfit122([]int{1, 2, 3, 4, 5}))
	// fmt.Println(maxProfit309([]int{1, 2, 3, 0, 2}))
	// fmt.Println(maxProfit714([]int{1, 2, 3, 8, 4, 9}, 2))
	// fmt.Println(maxProfit123([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	// fmt.Println(maxProfit188(2, []int{3, 2, 6, 5, 0, 3}))
	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	// fmt.Println(rob213([]int{1, 2, 3, 1}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
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

// lc 121 买卖股票的最佳时机,只交易一次
// 输入：[7,1,5,3,6,4]
// 输出：5
// 思路: dp的最简单版本,记录当前价格之前的最小值,同时记录当前价格的最大值,遍历后即可得
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}
	return maxProfit
}

// lc 122 买卖股票的最佳时机,可以交易很多次
func maxProfit122(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0] //交易第一天不买入为0,买入则成本为prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //要么没买观望,要么卖
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //要么持有观望,要么买
	}
	return dp[n-1][0]
}

// lc309. 买卖股票的最佳时机含冷冻期
func maxProfit309(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 { // base case i=0
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i-2 == -1 { // base case i=1
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i]) //第2天:要么第一天买过观望,要么第一天没买今天买
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) //解冻占一天
	}
	// fmt.Println(dp)
	return dp[n-1][0]
}

// 714. 买卖股票的最佳时机含手续费
func maxProfit714(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 { // base case i=0
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		if i-2 == -1 { // base case i=1
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i]-fee) //第2天:要么第一天买过观望,要么第一天没买今天买
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee) //买时记得扣费
	}
	// fmt.Println(dp)
	return dp[n-1][0]
}

// 123. 买卖股票的最佳时机 III 只允许交易两次
func maxProfit123(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxK := 2 // 最大可交易次数
	n := len(prices)
	dp := make([][][]int, n) //i表示天数,k表示交易次数,j表示持有状况:0表示未持有,1表示持有
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		for k := 0; k < maxK+1; k++ {
			dp[i][k] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := maxK; k > 0; k-- {
			if i-1 == -1 { // base case i=0
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i]) //注意频次需要遍历,同时买时频次减少1
		}
	}
	// 穷举了 n × maxK × 2 个状态
	return dp[n-1][maxK][0]
}

// lc188 你最多可以完成 k 笔交易
// 待优化点,如果k>=n/2,退化为无限次交易了,因为一次交易至少买入卖出各需要一次
// https://labuladong.github.io/algo/di-er-zhan-a01c6/yong-dong--63ceb/yi-ge-fang-3b01b/
func maxProfit188(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxK := k
	n := len(prices)
	dp := make([][][]int, n) //i表示天数,k表示交易次数,j表示持有状况:0表示未持有,1表示持有
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		for k := 0; k < maxK+1; k++ {
			dp[i][k] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := maxK; k > 0; k-- {
			if i-1 == -1 { // base case i=0
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i]) //注意频次需要遍历,同时买时频次减少1
		}
	}
	// 穷举了 n × maxK × 2 个状态
	return dp[n-1][maxK][0]
}

// lc42 接雨水 时间复杂度O(N2) 待优化
func trapOld(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans := 0
	for i := 1; i < len(height); i++ {
		leftMax, rightMax := 0, 0
		// 右边最高柱子
		for j := i; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}
		// 左边最高柱子
		for j := i; j >= 0; j-- {
			leftMax = max(leftMax, height[j])
		}
		// 计算能装的水
		ans += min(leftMax, rightMax) - height[i]
	}
	return ans
}

// lc42 接雨水 直接计算好每个位置的leftMax,rightMax
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans := 0
	n := len(height)
	leftMaxArr := make([]int, n) //备忘录
	rightMaxArr := make([]int, n)
	leftMaxArr[0] = height[0]
	rightMaxArr[n-1] = height[n-1]
	for i := 1; i < len(height); i++ {
		leftMaxArr[i] = max(leftMaxArr[i-1], height[i])
	}
	for j := n - 2; j >= 0; j-- {
		rightMaxArr[j] = max(rightMaxArr[j+1], height[j])
	}
	for i := 0; i < n; i++ {
		ans += min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}
	return ans
}

// lc198 打家劫舍
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	// base case
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])
	// dp[2] = max(dp[1], dp[0]+nums[2]) //找规律迭代
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	// fmt.Println(dp)
	return dp[n-1]
}

// lc198范围版本
func robRange(nums []int, start int, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start+1], nums[start])
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[end]
}

// lc213 打家劫舍2
func rob213(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//首尾房间不能同时被抢，那么只可能有三种不同情况：
	// 要么都不被抢[忽略]；要么第一间房子被抢最后一间不抢；要么最后一间房子被抢第一间不抢
	n := len(nums)
	firstRobRes := robRange(nums, 0, n-2)
	firstRobNotRes := robRange(nums, 1, n-1)
	return max(firstRobRes, firstRobNotRes)
}

// lc11 盛最多水的容器
// 间距在注定会缩小的情况下,另一个较小的先移动能找到最大的乘积值
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left, right := 0, len(height)-1
	var areaMax int
	for left <= right {
		area := min(height[left], height[right]) * (right - left)
		areaMax = max(areaMax, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return areaMax
}
