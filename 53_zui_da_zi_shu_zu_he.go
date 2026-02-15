/**
【题目53】最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

**/

package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := sol53_3(nums)
	fmt.Println(res)
}

func sol53_1(nums []int) int {
	// 动态规划找到 局部最优 & 全局最优
	maxLocal, maxGlobal := nums[0], nums[0]
	maxList := make([]int, 0)
	maxList = append(maxList, nums[0]) //
	for i := 1; i < len(nums); i++ {
		if nums[i]+maxList[i-1] > nums[i] {
			maxLocal = nums[i] + maxList[i-1]
		} else {
			maxLocal = nums[i]
		}
		maxList = append(maxList, maxLocal)
		if maxLocal > maxGlobal {
			maxGlobal = maxLocal
		}
	}
	return maxGlobal
}

func sol53_2(nums []int) int {
	globalMaxSum, currentMaxSum := -10000, 0
	for _, num := range nums {
		if currentMaxSum < 0 {
			currentMaxSum = num
		} else {
			currentMaxSum += num
		}
		// 更新最大和
		if currentMaxSum > globalMaxSum {
			globalMaxSum = currentMaxSum
		}
	}
	return globalMaxSum
}

func sol53_3(nums []int) int {
	// 局部最优 & 全局最优
	localMax, globalMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if localMax < 0 {
			localMax = nums[i]
		} else {
			localMax = localMax + nums[i]
		}
		if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}
