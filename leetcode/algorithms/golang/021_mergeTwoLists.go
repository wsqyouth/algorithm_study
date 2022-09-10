package main

// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"

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

// lc26题 思路1 k最小堆  思路2：分治算法
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	middle := length / 2
	leftList := mergeKLists(lists[:middle])
	rightList := mergeKLists(lists[middle:])
	return mergeTwoLists(leftList, rightList)
}

// lc160 相交链表的交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}
		if cur2 == nil {
			cur2 = headA
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}

//参考：https://blog.csdn.net/wo94chunjie/article/details/116274981
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0021.Merge-Two-Sorted-Lists/
