package main

import (
	"fmt"
	// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// removeElements 翻转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	var tempSlice []int
	for current != nil {
		tempSlice = append(tempSlice, current.Val)
		current = current.Next
	}
	dummyHead := &ListNode{}
	cur := dummyHead
	for i := len(tempSlice) - 1; i >= 0; i-- {
		cur.Next = &ListNode{Val: tempSlice[i]}
		cur = cur.Next
	}
	return dummyHead.Next
}

//reverseBetween 翻转指定区间的链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var preLeft *ListNode
	dummyHead := &ListNode{Next: head}
	preHead, current := dummyHead, head //preHead方便找左边界前驱
	count := 0
	var isRange bool
	var tempSlice []int
	for current != nil {
		if count == left-1 {
			preLeft = preHead
			isRange = true
		}
		if isRange {
			tempSlice = append(tempSlice, current.Val)
		}
		if count == right-1 {
			break
		}
		count++
		preHead = preHead.Next
		current = current.Next
	}

	//区间翻转链表
	tempHead := &ListNode{}
	tempCur := tempHead
	for i := len(tempSlice) - 1; i >= 0; i-- {
		tempCur.Next = &ListNode{Val: tempSlice[i]}
		tempCur = tempCur.Next
	}

	// 组装
	tempCur.Next = current.Next
	preLeft.Next = tempHead.Next
	return dummyHead.Next
}

// 删除重复的元素[保留一个] 11233  => 123
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: head.Val}
	slow, fast := newHead, head
	for fast != nil {
		if fast.Val != slow.Val {
			fmt.Println(fast.Val, slow.Val)
			slow.Next = &ListNode{Val: fast.Val}
			slow = slow.Next
		}
		fast = fast.Next
	}
	return newHead
}

// 删除重复的元素[一个也不留] 11123  => 23
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Next: head} //若重复仍去掉
	newCurrent := dummyHead
	count := 1

	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val == slow.Val {
			count++
		} else {
			fmt.Println(fast.Val, slow.Val)
			if count > 1 {
				newCurrent.Next = fast //妙
			} else {
				newCurrent = newCurrent.Next
			}
			slow = fast
			count = 1
		}
		fast = fast.Next
	}
	if count > 1 {
		newCurrent.Next = nil
	}
	return dummyHead.Next
}
