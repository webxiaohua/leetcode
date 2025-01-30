/*
【题目102】二叉树的层序遍历
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
		3
	  /   \
	9		20
           / \
  	      15  7

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/

package main

import "fmt"

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

func main() {
	root := &TreeNode102{
		Val: 3,
		Left: &TreeNode102{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode102{
			Val:   20,
			Left:  &TreeNode102{Val: 15},
			Right: &TreeNode102{Val: 7},
		},
	}
	res := sol102_1(root)
	fmt.Println(res)
}

// BFS + 队列，把当前层的节点加入队列，遍历队列，把下一层级的元素加入队列，同时剔除掉当前层级的元素
func sol102_1(root *TreeNode102) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode102{root}
	for {
		if len(queue) == 0 {
			break
		}
		var tmpArr []int
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			tmpArr = append(tmpArr, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[cnt:]
		res = append(res, tmpArr)
	}
	return res
}
