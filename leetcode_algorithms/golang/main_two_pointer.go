package main

import (
	"fmt"
	"sort"
)

/*
func main() {
	sumTest()
}
*/
func sumTest() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	fmt.Println(twoSumNew(nums, 9))
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

// lc1 两数之和
func twoSum(nums []int, target int) []int {
	var targetIndex []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				targetIndex = append(targetIndex, i, j)
				return targetIndex
			}
		}
	}
	return nil
}

// 将数据和索引记录下来，遍历时检查是否在map中
func twoSumNew(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// lc15 三数之和
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	length := len(nums)
	sum := 0
	for index := 1; index < length-1; index++ {
		start, end := 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && index < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum = nums[start] + nums[index] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}

	}
	return result
}
