/**
【题目42】接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
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
	res := sol42_1(height)
	fmt.Println(res)
}

func sol42_1(height []int) int {
	res := 0
	if len(height) < 3 {
		return res
	}
	maxLeft, maxRight, left, right := 0, 0, 0, len(height)-1
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

func getMax(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
