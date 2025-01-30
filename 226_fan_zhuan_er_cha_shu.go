/*
【题目226】翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

示例 1：
输入：root = [4,2,7,1,3,6,9]
		4
	  /   \
	2		7
  /   \    / \
 1	   3  6	  9
输出：[4,7,2,9,6,3,1]
		4
	  /   \
	7		2
  /   \    / \
 9	   6  3	  1

示例 2：
输入：root = [2,1,3]
输出：[2,3,1]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目范围在 [0, 100] 内
-100 <= Node.val <= 100
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	res := sol226_1(node)
	fmt.Println(res)
}

func sol226_1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	sol226_1(root.Left)
	sol226_1(root.Right)
	return root
}
