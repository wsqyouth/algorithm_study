package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums:=[]int{-1,0,1,2,-1,-4}
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	fmt.Println(threeSum(nums))
}

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

	//fmt.Println(nums)
	return result
}
func threeSum1(nums []int) [][]int {
	var target [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					var targetTemp []int
					targetTemp = append(targetTemp, nums[i], nums[j], nums[k])
					fmt.Println(targetTemp, checkExist(target, targetTemp))
					if !checkExist(target, targetTemp) {
						target = append(target, targetTemp)
					}
				}
			}
		}
	}
	return target
}

// 检查一维数组是否存在于二维数组中
func checkExist(target [][]int, targetTemp []int) bool {
	for _, arr := range target {
		m := make(map[int]bool)
		for i := 0; i < len(arr); i++ {
			m[arr[i]] = true
		}
		var count int
		for _, val := range targetTemp {
			if _, ok := m[val]; ok {
				count++
			}
		}
		if count == len(targetTemp) {
			return true
		}
	}
	return false
}

//总结:
//自己想的办法：数组去重，但是这种办法依赖于判断两个切片是否一致，自己实现的方法有漏洞，针对重复元素使用map判断是错误的
//参考网上是思路，使用双指针法实现，后续继续研究下
//2022年04月03日19:14
