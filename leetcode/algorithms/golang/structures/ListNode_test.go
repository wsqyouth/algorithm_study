package structures

import (
	"reflect"
	"testing"
)

func TestInts2List(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *ListNode
	}{
		{
			name:"ints to list",
			nums:[]int{1, 2, 3}
			want:&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints2List(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints2List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList2Ints(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := List2Ints(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List2Ints() = %v, want %v", got, tt.want)
			}
		})
	}
}
