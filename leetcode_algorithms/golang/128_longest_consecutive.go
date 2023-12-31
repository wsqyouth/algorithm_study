package main

import "fmt"

func main() {
	num := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(num))
}

// lc128 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 思路: 类似公司员工年龄基数统计的变形
func longestConsecutive(nums []int) int {
	existMap := make(map[int]bool)
	for _, num := range nums {
		existMap[num] = true
	}

	res := 0
	for num := range existMap {
		if existMap[num-1] { //num不是第一个首字符则忽略
			continue
		}
		curNum := num
		curLen := 1
		for existMap[curNum+1] { //写的好
			curNum++
			curLen++
		}
		res = max(res, curLen)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
