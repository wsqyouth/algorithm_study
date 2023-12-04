package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		newLists := []*ListNode{}

		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				newLists = append(newLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				newLists = append(newLists, lists[i])
			}
		}

		lists = newLists
	}

	return lists[0]
}

func main() {
	// 创建测试数据
	lists := []*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{7, nil}}},
		&ListNode{2, &ListNode{5, &ListNode{8, nil}}},
		&ListNode{3, &ListNode{6, &ListNode{9, nil}}},
	}
	/*
		// 合并链表
		result := mergeKLists(lists)

		// 打印结果
		for node := result; node != nil; node = node.Next {
			fmt.Print(node.Val, " ")
		}
	*/
	// ------------------------- 堆实现 ----------------------------------------

	// 合并链表
	resultHeap := mergeKListsHeap(lists)
	// 打印结果
	for node := resultHeap; node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}

}

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		curr.Next = minNode
		curr = curr.Next

		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummy.Next
}

/*
首先将所有链表的头节点添加到最小堆中。然后，它从堆中取出最小的节点，并将该节点的下一个节点添加到堆中。
这个过程一直持续到堆为空。每次从堆中取出一个节点时，都将其添加到结果链表的末尾。
*/
