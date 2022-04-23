package structures

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next ListNode
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	list := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return list.Next
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制为100,超过则检查是否有环或者调整
	limit := 100
	count := 0

	res := []int{}
	for head != nil {
		count++
		if count > limit {
			msg := fmt.Sprintf("链条深度超过%d,可能出现环,请检查错误或者放宽限制", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
