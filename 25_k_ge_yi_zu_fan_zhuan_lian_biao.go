/**
【题目25】K个一组翻转链表
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

提示：
链表中的节点数目为 n
1 <= k <= n <= 5000
0 <= Node.val <= 1000

进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
**/

package main

import "fmt"

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

func main() {
	node := &ListNode25{Val: 1, Next: &ListNode25{Val: 2, Next: &ListNode25{Val: 3, Next: &ListNode25{Val: 4, Next: &ListNode25{Val: 5}}}}}
	res := sol25_1(node, 2)
	for {
		if res != nil {
			fmt.Println(res.Val)
			res = res.Next
		} else {
			break
		}
	}
}

func sol25_1(head *ListNode25, k int) *ListNode25 {
	res := &ListNode25{Next: head}
	prev := res
	for {
		var nodes []*ListNode25
		tmpNode := prev
		for i := 0; i < k; i++ {
			if tmpNode.Next != nil {
				nodes = append(nodes, tmpNode.Next)
				tmpNode = tmpNode.Next
			}
		}
		if len(nodes) != k {
			break
		}
		// nodes 反转链表
		nodes[0].Next = nodes[k-1].Next
		for i := k - 1; i > 0; i-- {
			nodes[i].Next = nodes[i-1]
		}
		prev.Next = nodes[k-1]
		prev = nodes[0]
	}
	return res.Next
}
