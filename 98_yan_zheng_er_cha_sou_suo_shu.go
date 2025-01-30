/*
【题目98】验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]
输出：true
		2
	  /   \
	1		3

示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
		5
	  /   \
	1		4
           / \
  	      3   6
解释：根节点的值是 5 ，但是右子节点的值是 4 。

提示：
树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode98{
		Val: 5,
		Left: &TreeNode98{
			Val: 4,
		},
		Right: &TreeNode98{
			Val: 6,
			Left: &TreeNode98{
				Val: 3,
			},
			Right: &TreeNode98{
				Val: 7,
			},
		},
	}
	res := sol98_1(root)
	fmt.Println(res)
}

type TreeNode98 struct {
	Val   int
	Left  *TreeNode98
	Right *TreeNode98
}

// 当前节点对比上下限，每次递归都更新上下限
func sol98_1(root *TreeNode98) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(root *TreeNode98, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}
