/**
【题目23】合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
**/

package main

type ListNode23 struct {
	Val  int
	Next *ListNode23
}

func main() {
	node1 := &ListNode23{Val: 1, Next: &ListNode23{Val: 3, Next: &ListNode23{Val: 5}}}
	node2 := &ListNode23{Val: 2, Next: &ListNode23{Val: 4, Next: &ListNode23{Val: 7}}}
	node3 := &ListNode23{Val: 1, Next: &ListNode23{Val: 2, Next: &ListNode23{Val: 6}}}
	lists := []*ListNode23{node1, node2, node3}
	res := sol23_1(lists)
	for {
		if res != nil {
			println(res.Val)
			res = res.Next
		} else {
			break
		}
	}
}

// 执行遍历，每次找到最小值取出拼接到返回结果中，直到 lists 为空
func sol23_1(lists []*ListNode23) *ListNode23 {
	res := &ListNode23{}
	pointer := res
	for {
		// 找到最小值
		if len(lists) == 0 {
			break
		}
		minNode := lists[0]
		if minNode == nil {
			lists = lists[1:]
			continue
		}
		minIdx := 0
		for i, node := range lists {
			if node == nil {
				continue
			}
			if minNode.Val > node.Val {
				minNode = node
				minIdx = i
			}
		}
		pointer.Next = minNode
		pointer = pointer.Next
		if minNode.Next == nil {
			lists = append(lists[:minIdx], lists[minIdx+1:]...)
		} else {
			lists[minIdx] = lists[minIdx].Next
		}
	}
	return res.Next
}
