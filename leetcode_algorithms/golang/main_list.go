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

// lc2 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	carry := 0
	cur := dummyNode
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return dummyNode.Next
}

// lc148. 排序链表
// 这道题还有迭代归并法,后续看,非常值得学习 https://leetcode.cn/problems/sort-list/solutions/2358097/ru-he-zheng-ming-pai-xu-lian-biao-shi-ku-tf9m/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 寻找中间节点
	fast, slow := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	// 拆分两个链表
	prev.Next = nil

	// 继续拆分排序
	left := sortList(head)
	right := sortList(slow)

	// 合并有序链表
	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}
