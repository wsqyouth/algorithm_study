package main

import (
	"fmt"

	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

//func main() {
//mergeKListsTest()
//}
func mergeKListsTest() {
	list := structures.Ints2List([]int{1, 4, 5})
	list1 := structures.Ints2List([]int{1, 3, 4})
	list2 := structures.Ints2List([]int{2, 6})
	lists := []*structures.ListNode{list, list1, list2}
	structures.PrintList(mergeKLists(lists))
}
func partitionTest() {
	list := structures.Ints2List([]int{1, 4, 3, 2, 5, 2})
	structures.PrintList(partition(list, 3))
}

func hasCycleTest() {
	list := structures.Ints2List([]int{3, 2, 0, -4})
	// 构造有环链表,禁止打印
	fmt.Println(hasCycle(list))
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

func oddEvenListTest() {
	list := structures.Ints2List([]int{2, 1, 3, 5, 6, 4, 7})
	resList := oddEvenList(list)
	structures.PrintList(resList)
}

/*
func deleteNodeTest() {
	list := structures.Ints2List([]int{4, 5, 1, 9})
	current := list
	var deleteNodePointer *structures.ListNode
	for current != nil {
		if current.Val == 5 {
			deleteNodePointer = current
			break
		}
		current = current.Next
	}
	deleteNode(deleteNodePointer)
	structures.PrintList(list)
}
*/
func removeElementsTest() {
	list := structures.Ints2List([]int{7, 7, 7, 7})
	resList := removeElements(list, 7)
	structures.PrintList(resList)
}
func reverseListTest() {
	list := structures.Ints2List([]int{1, 2})
	resList := reverseList(list)
	structures.PrintList(resList)
}

func reverseBetweenTest() {
	list := structures.Ints2List([]int{1, 2, 3, 4, 5})
	resList := reverseBetween(list, 2, 4)
	structures.PrintList(resList)
}

func deleteDuplicatesTest() {
	list := structures.Ints2List([]int{1, 1})
	resList := deleteDuplicatesII(list)
	structures.PrintList(resList)
}

func mergeTwoListsTest() {
	list1 := structures.Ints2List([]int{1, 1, 2})
	list2 := structures.Ints2List([]int{2, 3, 4, 5})

	resList := mergeTwoLists(list1, list2)
	structures.PrintList(resList)
}
