/**
【题目24】两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

提示：
链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
**/

package main

import "fmt"

type ListNode24 struct {
	Val  int
	Next *ListNode24
}

func main() {
	head := &ListNode24{
		Val: 1,
		Next: &ListNode24{
			Val:  2,
			Next: nil,
		},
	}
	res := sol24_2(head)
	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

func sol24_1(head *ListNode24) *ListNode24 {
	res := &ListNode24{}
	res.Next = head
	prev := res
	for {
		if prev.Next == nil || prev.Next.Next == nil {
			break
		}
		node1 := prev.Next
		node2 := prev.Next.Next
		// 交换两个节点
		prev.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		// 移动prev
		prev = node1
	}
	return res.Next
}

func sol24_2(head *ListNode24) *ListNode24 {
	res := &ListNode24{}
	res.Next = head
	prev := res
	for {
		if prev.Next == nil || prev.Next.Next == nil {
			break
		}
		node1, node2 := prev.Next, prev.Next.Next
		prev.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		prev = node1
	}
	return res.Next
}
