package main

import (
	"fmt"
	"github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// removeNthFromEnd 快慢指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head //秒
	for fast != nil {
		if n <= 0 { //快指针先走了n步后,慢指针开始走,同时记录慢指针的上一个节点
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	slow.Next = nil
	return dummyHead.Next
}

// removeNthFromEnd1 先记录总数,之后重新记录
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	len := 0
	current := head
	for current != nil {
		len++
		current = current.Next
	}
	fmt.Println(len)
	if n == len {
		deleteNode := head
		head = head.Next
		deleteNode.Next = nil
		return head
	}
	i := 0
	current = head
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}
