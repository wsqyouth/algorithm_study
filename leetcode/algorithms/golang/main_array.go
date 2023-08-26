package main

import "fmt"

func main() {
	// mergeInplaceTest()
	// fmt.Println(search([]int{1}, 3))
	// fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	// fmt.Println(search2([]int{1, 0, 1, 1, 1}, 0))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}

//
func mergeInplaceTest() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

// lc88 合并两个有序数组,nums1开辟空间长度为二者之和
// 从后向前处理，无需移动数据
func merge(nums1 []int, m int, nums2 []int, n int) {
	var curVal int
	for nums1Index, nums2Index, index := m-1, n-1, m+n-1; nums1Index >= 0 || nums2Index >= 0; index-- {
		if nums1Index == -1 {
			curVal = nums2[nums2Index]
			nums2Index--
		} else if nums2Index == -1 {
			curVal = nums1[nums1Index]
			nums1Index--
		} else if nums1[nums1Index] > nums2[nums2Index] {
			curVal = nums1[nums1Index]
			nums1Index--
		} else {
			curVal = nums2[nums2Index]
			nums2Index--

		}
		nums1[index] = curVal
	}
}

// lc287 寻找重复数
// nums = [1,3,4,2,2]
func findDuplicate(nums []int) int {
	slow := 0
	fast := nums[slow]
	// fast 向后移动两步，slow 移动一步
	for slow != fast {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	// 相遇后,fast从下一个位置开始,slow从零开始
	fast = nums[fast]
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// lc33 搜索排序旋转数组内目标值的索引  target = 3
// [4,5,6,] 7,0,1,2 左边的是有序的,且target位于[nums[left],nums[mid-1]],则收缩范围到左边
// 7,0,1,2,[3,4,5,6] 右边是有序的,且target位于[nums[mid+1],nums[right]],则收缩范围到右边
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		// fmt.Println(mid, nums[mid])
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// lc81 搜索旋转数组带重复元素
func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] >= nums[left] { //前部分有序
			if target >= nums[left] && target < nums[mid] { //前半部分升序
				right = mid - 1
			} else {
				left = mid + 1 //后半部分无序
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

//lc32 寻找target的左边界,右边界
// nums = [5,7,7,8,8,10], target = 8
func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 别返回，锁定左侧边界
			right = mid - 1
		}
	}
	// 判断 target 是否存在于 nums 中
	if left < 0 || left >= len(nums) {
		return -1
	}
	// 判断一下 nums[left] 是不是 target
	if nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 别返回，锁定右侧边界
			left = mid + 1
		}
	}

	if right < 0 || right >= len(nums) {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}

// lc80
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
func removeDuplicatesBad(nums []int) int { //这是我写的，太复杂了，虽然是双指针也有漏洞
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	count := 1
	for fast < len(nums)-1 {
		for fast < len(nums)-1 && nums[fast] == nums[fast+1] {
			count++
			fast++
		}
		fmt.Println(fast, nums[fast], count)

		if count > 2 {
			slow += 2
		} else {
			slow += count
		}
		count = 1
		fast++
	}
	return slow
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	count := 0
	for fast < len(nums) {
		if count < 2 || nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
		fmt.Println(fast, nums[fast], count)
		fast++
	}
	return slow
}
