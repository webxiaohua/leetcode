/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807

输入：l1 = [0], l2 = [0]
输出：[0]

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
**/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	res := &ListNode2{Next: &ListNode2{}}
	tmp := res.Next
	var v1, v2, carry int
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode2{
			Val: (v1 + v2 + carry) % 10,
		}
		carry = (v1 + v2 + carry) / 10
		tmp = tmp.Next
	}
	return res.Next.Next
}

func main() {
	l1 := &ListNode2{Val: 2, Next: &ListNode2{
		Val: 4,
		Next: &ListNode2{
			Val:  3,
			Next: nil,
		},
	}}
	l2 := &ListNode2{Val: 5, Next: &ListNode2{
		Val: 6,
		Next: &ListNode2{
			Val:  4,
			Next: nil,
		},
	}}
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current, carry := res, 0
	for l1 != nil || l2 != nil || carry != 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return res.Next
}
*/
