/**
【题目42】接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例1:
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例2:
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
**/

package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := sol_42_2(height)
	fmt.Println(res)
}

func sol_42_1(height []int) int {
	// 计算每一个格子可以容纳的雨水量，需要知道3个要素：  左侧的最高点、右侧的最高点、当前的高度
	// 按照短板理论，取左右侧最高点的较小值，减去当前高度，即为本格所能容纳的雨水量
	// 如果左侧高点为较小值，计算完成后左侧指针往右位移
	// 如果右侧高点为较小值，计算完成后右侧指针往左位移
	// 两者相遇时退出
	if len(height) < 3 {
		// 不足3个元素直接返回0
		return 0
	}
	left, right := 0, len(height)-1
	var maxLeft, maxRight, res int
	for left < right {
		maxLeft = getMax(height[left], maxLeft)
		maxRight = getMax(height[right], maxRight)
		if maxLeft <= maxRight {
			res += maxLeft - height[left]
			left++
		} else {
			res += maxRight - height[right]
			right--
		}
	}
	return res
}

func getMax(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func sol_42_2(height []int) int {
	if len(height) < 3 {
		return 0
	}
	maxLeft, maxRight, left, right, res := 0, 0, 0, len(height)-1, 0
	for left < right {
		maxLeft = getMax(maxLeft, height[left])
		maxRight = getMax(maxRight, height[right])
		if maxLeft < maxRight {
			res += maxLeft - height[left]
			left++
		} else {
			res += maxRight - height[right]
			right--
		}
	}
	return res
}
