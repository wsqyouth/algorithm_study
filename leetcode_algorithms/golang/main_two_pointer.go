package main

import (
	"fmt"
	"sort"
)

func main() {
	sumTest()
}

func sumTest() {
	// nums := []int{2, 7, 11, 15}
	// fmt.Println(twoSum(nums, 9))
	// fmt.Println(twoSumNew(nums, 9))
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumNew(nums))
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

func threeSumNew(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	// i left right
	for i := 0; i < n-2; i++ {
		// 对 i 进行去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 对left和right进行去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right-- // 和太大，右指针左移减小
			} else {
				left++ // 和太小，左指针右移增大
			}
		}
	}
	return res
}

func threeSumClosestNew(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[n-1]

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right--
			} else {
				left++
			}
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}
		}
	}
	return closestSum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
首先对输入的数组进行了排序，然后通过固定一个数，移动两个指针来找到最接近目标值的三个数的和。如果当前的和大于目标值，就将右指针向左移动；如果当前的和小于目标值，就将左指针向右移动。在每次计算和的时候，都会检查当前的和是否比之前找到的更接近目标值，如果是，就更新结果。
*/
