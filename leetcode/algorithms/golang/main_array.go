package main

import "fmt"

func main() {

	mergeInplaceTest()
}

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
