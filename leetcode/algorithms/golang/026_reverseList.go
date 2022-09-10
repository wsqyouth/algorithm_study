package main

import (
	"fmt"
	// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// removeElements 翻转链表,lc 206题
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

/*
141:判断链表是否有环
每次移动慢指针一步，而移动快指针两步。每一次迭代，快速指针将额外移动一步。
如果环的长度为 M，经过 M 次迭代后，快指针肯定会多绕环一周，并赶上慢指针。
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// lc142 链表入环的第一个结点
func detectCycle(head *ListNode) *ListNode {
	var hasCycle bool
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	// 慢指针从头开始走
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

//输入：head = [1,4,3,2,5,2], x = 3 输出：[1,2,2,4,3,5]
// leetcdoe 86题
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	smallListHead, largeListHead := &ListNode{}, &ListNode{}
	smallListCur, largeListCur := smallListHead, largeListHead
	current := head
	for current != nil {
		if current.Val < x {
			smallListCur.Next = current
			smallListCur = smallListCur.Next
		} else {
			largeListCur.Next = current
			largeListCur = largeListCur.Next
		}
		// current需断掉原有的连接关系，同时向前移动
		tempNode := current.Next
		current.Next = nil
		current = tempNode
	}
	smallListCur.Next = largeListHead.Next
	return smallListHead.Next
}
