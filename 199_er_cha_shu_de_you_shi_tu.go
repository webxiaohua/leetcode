/**
【题目199】二叉树的右视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]
       [1]
	  /   \
	2	   [3]
     \       \
  	  5      [4]

示例 2：
输入：root = [1,2,3,4,null,null,null,5]
输出：[1,3,4,5]
       [1]
	  /   \
	2	   [3]
   /
  [4]
 /
[5]

示例3：
输入：root = [1,null,3]
输出：[1,3]

示例4：
输入：root = []
输出：[]

提示：
二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
**/

package main

import "fmt"

type TreeNode199 struct {
	Val   int
	Left  *TreeNode199
	Right *TreeNode199
}

func main() {
	root := &TreeNode199{
		Val: 1,
		Left: &TreeNode199{
			Val: 2,
			Left: &TreeNode199{
				Val: 4,
				Left: &TreeNode199{
					Val: 5,
				},
			},
		}, Right: &TreeNode199{
			Val: 3,
		},
	}
	res := sol199_1(root)
	fmt.Println(res)
}

func sol199_1(root *TreeNode199) []int {
	// 层序遍历，取最右侧数据
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode199{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
