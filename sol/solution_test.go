package sol

import (
	"reflect"
	"testing"
)

func CreateNumList(nums *[]int) *ListNode {
	var result, cur *ListNode
	arr := *nums
	for idx, val := range arr {
		if idx == 0 {
			result = &ListNode{Val: val}
			cur = result
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return result
}
func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "l1 = [2,4,3], l2 = [5,6,4]",
			args: args{l1: CreateNumList(&[]int{2, 4, 3}), l2: CreateNumList(&[]int{5, 6, 4})},
			want: CreateNumList(&[]int{7, 0, 8}),
		},
		{
			name: "l1 = [0], l2 = [0]",
			args: args{l1: CreateNumList(&[]int{0}), l2: CreateNumList(&[]int{0})},
			want: CreateNumList(&[]int{0}),
		},
		{
			name: "l1 = [0], l2 = [0]",
			args: args{l1: CreateNumList(&[]int{0}), l2: CreateNumList(&[]int{0})},
			want: CreateNumList(&[]int{0}),
		},
		{
			name: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]",
			args: args{l1: CreateNumList(&[]int{9, 9, 9, 9, 9, 9, 9}), l2: CreateNumList(&[]int{9, 9, 9, 9})},
			want: CreateNumList(&[]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
