package main

import (
	"fmt"
	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

func main() {
	middleNodeTest()
}

func twoSumTest() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
	fmt.Println(twoSumNew(nums, 6))
}

func removeNthFromEndTest() {
	list := structures.Ints2List([]int{1, 2, 3, 4, 5})
	removeList := removeNthFromEnd(list, 2)
	structures.PrintList(removeList)
}

func middleNodeTest() {
	list := structures.Ints2List([]int{1, 2})
	middList := middleNode(list)
	structures.PrintList(middList)
}
