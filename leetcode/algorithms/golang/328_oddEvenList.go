package main

import (
// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// oddEvenList 按照索引奇数偶数分为两个链表，最后连接起来
func oddEvenList(head *ListNode) *ListNode {
	oddListHead, evenListHead := &ListNode{}, &ListNode{}
	oddList, evenList := oddListHead, evenListHead
	current := head
	i := 1
	for current != nil {
		if i%2 != 0 {
			oddList.Next = &ListNode{Val: current.Val}
			oddList = oddList.Next
		} else {
			evenList.Next = &ListNode{Val: current.Val}
			evenList = evenList.Next
		}
		i++
		current = current.Next
	}
	oddList.Next = evenListHead.Next
	return oddListHead.Next
}
