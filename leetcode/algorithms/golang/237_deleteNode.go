package main

import (
// "fmt"
// "github.com/wsqyouth/algorithm_study/leetcode/algorithms/golang/structures"
)

// ListNode define
// type ListNode = structures.ListNode

// deleteNode 删除指定节点，直接将下一节点覆盖即可
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
