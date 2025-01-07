/**
【题目45】跳跃游戏II
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2

提示:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]
**/

package main

import "fmt"

func main() {
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	res := sol45_1(nums)
	fmt.Println(res)
}

func sol45_1(nums []int) int {
	minStep := 0
	if len(nums) < 2 {
		return minStep
	}
	// 当前可达最大位置
	farthest := 0
	// 当前边界位置
	currentEnd := 0
	for i := range nums {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		// 当触碰到边界时需要更新边界为最大可达位置，并且步数加一
		if i == currentEnd {
			minStep++
			currentEnd = farthest
			// 当边界超过数组长度时，直接返回
			if currentEnd >= len(nums)-1 {
				break
			}
		}
	}
	return minStep
}

/*
[2,3,1,1,4]
i=0, farthest = 2, currentEnd = 2, minStep = 1
i=1, farthest = 4, currentEnd = 2, minStep = 1
i=2, farthest = 4, currentEnd = 4, minStep = 2
*/
