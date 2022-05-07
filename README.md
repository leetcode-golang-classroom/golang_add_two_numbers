# golang_add_two_numbers

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]

```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

```

**Constraints:**

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

## 解析

題目把數字以鏈結串列的方式表示把 digit 用到序的方式放到每個結點

要求實作兩個數字相加之後的表示式

首先要知道數字化成表示是特性舉例來說 342 的表示是就是 [2,4,3]

而加法恰巧是可以從個位數開始加

當digit 與前一位數的 carry 相加超過 base 10 就要進位

初始化 carry := 0 , digit := 0

loop 兩個 linked list

每次 digit := (carry + digit_1+ digit_2)%10

        carry =  carry + digit_1+ digit_2 ≥ 10 ? 1: 0

以最長的數字為基礎需要檢查 M + 1 次確認是否有進位

所時間複雜度 O(N)

## 程式碼

```go
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
	for l1!= nil || l2!= nil {
	 digit_1, digit_2 := 0, 0
	 if l1 != nil {
	   digit_1 = l1.Val
	   l1 = l1.Next
	 }
	 if l2 != nil {
	   digit_2 = l2.Val
	   l2 = l2.Next
	 }
	 sum := carry + digit_1 + digit_2
	 if sum >= 10 {
	   carry = 1
	 } else {
	   carry = 0
	 }
	 digit := sum % 10
	     if resultSum == nil {
	    resultSum = &ListNode{Val: digit}
	    cur = resultSum
	 }	else {
	    cur.Next = &ListNode{Val: digit}
	    cur = cur.Next
	 }
	}
	if carry != 0 {
	    cur.Next = &ListNode{Val: 1}
	} 
	return resultSum
}
```

## 困難點

1. 需要注意到 carry 

## Solve Point

- [x]  Understand what the problem would like to solve
- [x]  Analysis Complexity