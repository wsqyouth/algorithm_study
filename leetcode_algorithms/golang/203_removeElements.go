package main

import (
// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// removeElements 去除链表中指定元素,记下重复元素的前驱preSlow,对重复元素进行去重
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, fast := dummyHead, head //秒
	for fast != nil {
		if fast.Val != val {
			preSlow = preSlow.Next
			fast = fast.Next
		} else {
			for fast != nil && fast.Val == val {
				fast = fast.Next
			}
			preSlow.Next = fast
		}
	}
	return dummyHead.Next
}
