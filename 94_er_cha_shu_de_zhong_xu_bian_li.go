/*
【题目94】二叉树de中序遍历
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶:递归算法很简单，你可以通过迭代算法完成吗？
*/
package main

import "fmt"

func main() {
	node := &TreeNode94{Val: 1, Right: &TreeNode94{Val: 2, Left: &TreeNode94{Val: 3}}}
	res := sol94_2(node)
	fmt.Println(res)
}

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

// 递归法 - 左，根，右
func sol94_1(root *TreeNode94) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, sol94_1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, sol94_1(root.Right)...)
	return res
}

func sol94_2(root *TreeNode94) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, sol94_2(root.Left)...)
	res = append(res, root.Val)
	res = append(res, sol94_2(root.Right)...)
	return res
}
