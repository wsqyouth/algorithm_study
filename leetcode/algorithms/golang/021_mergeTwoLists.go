package main

import (
// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// deleteNode 合并双链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummpHead := &ListNode{} //借用一个数字做比较,返回时丢弃
	newCurrent, cur1, cur2 := dummpHead, list1, list2
	for cur1 != nil && cur2 != nil {
		// fmt.Println(cur1.Val, cur2.Val)
		if cur1.Val <= cur2.Val {
			newCurrent.Next = cur1
			cur1 = cur1.Next
		} else {
			newCurrent.Next = cur2
			cur2 = cur2.Next
		}

		newCurrent = newCurrent.Next
	}
	if cur1 == nil {
		newCurrent.Next = cur2
	}
	if cur2 == nil {
		newCurrent.Next = cur1
	}

	return dummpHead.Next
}

//参考：https://blog.csdn.net/wo94chunjie/article/details/116274981
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0021.Merge-Two-Sorted-Lists/
