/**
【题目238】除自身以外数组的乘积
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:
输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：
2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
**/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := sol238_2(nums)
	fmt.Println(res)
}

func sol238_1(nums []int) []int {
	res := make([]int, len(nums))
	// 采用两个数组分别存放 i 左侧的乘积和 i 右侧的乘积，再遍历一次计算
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func sol238_2(nums []int) []int {
	res := make([]int, len(nums))
	// 不引入额外空间，直接使用 res，两次遍历，一次从左往右，计算左侧乘积，另一侧从右往左计算乘积
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	// 分别乘以右侧数据
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res[i] = res[i] * nums[j]
		}
	}
	return res
}
