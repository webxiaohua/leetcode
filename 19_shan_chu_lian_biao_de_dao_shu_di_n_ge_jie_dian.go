/**
【题目19】删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？
**/

package main

import "fmt"

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

func main() {
	head := &ListNode19{
		Val: 1, Next: &ListNode19{
			Val: 2,
			Next: &ListNode19{
				Val: 3,
				Next: &ListNode19{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	res := sol19_1(head, 1)
	for {
		if res == nil {
			return
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

func sol19_1(head *ListNode19, n int) *ListNode19 {
	// 双指针， 第一个指针先前进n步，然后两个指针一起前进，当第一个指针到达尾部时第二个指针刚好到达倒数的第k个元素的前一个元素
	var fast, slow *ListNode19
	fast = head
	cnt := 1
	for {
		if fast.Next == nil {
			if slow != nil {
				slow.Next = slow.Next.Next
			} else {
				head = head.Next
			}
			break
		}
		if cnt == n {
			slow = head
		} else if cnt > n {
			slow = slow.Next
		}
		fast = fast.Next
		cnt++
	}
	return head
}
