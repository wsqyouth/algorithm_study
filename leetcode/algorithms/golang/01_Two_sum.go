package main

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
