package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(nums []int, begin, end int) int {
	pivot := nums[begin]
	for begin < end {
		for begin < end && nums[end] >= pivot {
			end--
		}
		nums[begin] = nums[end]
		for begin < end && nums[begin] <= pivot {
			begin++
		}
		nums[end] = nums[begin]
	}
	nums[begin] = pivot
	return begin
}

func quickSort(nums []int, begin, end int) {
	if begin < end {
		pivot := partition(nums, begin, end)
		quickSort(nums, begin, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func findKthLargest(nums []int, k int) int {
	begin, end := 0, len(nums)-1
	target := len(nums) - k
	for {
		pivot := partition(nums, begin, end)
		if pivot == target {
			return nums[pivot]
		} else if pivot < target {
			begin = pivot + 1
		} else {
			end = pivot - 1
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := []int{3, 2, 1, 5, 6, 4}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums) // 输出排序后的数组

	nums = []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k)) // 输出第k大的元素
}

//quickSort函数实现了快速排序，findKthLargest函数实现了获取第k大的元素。这两个函数都使用了partition函数来进行分区。
