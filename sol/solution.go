package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultSum *ListNode
	var cur *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		digit_1, digit_2 := 0, 0
		if l1 != nil {
			digit_1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			digit_2 = l2.Val
			l2 = l2.Next
		}
		sum := digit_1 + digit_2 + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		digit := sum % 10
		if resultSum == nil {
			resultSum = &ListNode{Val: digit}
			cur = resultSum
		} else {
			cur.Next = &ListNode{Val: digit}
			cur = cur.Next
		}
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: 1}
	}
	return resultSum
}
