package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	start := len(nums) - k
	var numsLast []int
	for i := 0; i < k; i++ {
		numsLast = append(numsLast, nums[start+i])
	}
	for i := start - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < start-1; i++ {
		nums[i] = numsLast[i]
	}
}

// @lc code=end
