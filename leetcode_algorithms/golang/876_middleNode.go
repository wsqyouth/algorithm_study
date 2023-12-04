package main

import (
// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// middleNode1 快慢指针法
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	len := 0
	current := head
	for current != nil {
		len++
		current = current.Next
	}
	if len%2 == 0 {
		return slow.Next
	}
	return slow
}

// middleNode1 先记录总数,之后重新记录,记得考虑单节点
func middleNode1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	len := 0
	current := head
	for current != nil {
		len++
		current = current.Next
	}
	// fmt.Println(len)
	middle := len / 2
	i := 0
	current = head
	for current != nil {
		if i == middle-1 {
			return current.Next
		}
		i++
		current = current.Next
	}
	return nil
}
