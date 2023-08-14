package main

import "fmt"

func main() {
	// mergeInplaceTest()
	fmt.Println(search([]int{1}, 3))
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
