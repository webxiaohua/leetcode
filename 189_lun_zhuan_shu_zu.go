/**
【题目189】轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：
1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 10^5

进阶：
尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
**/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sol189_2(nums, 3)
	fmt.Println(nums)
}

func sol189_1(nums []int, k int) {
	// 复制一个新切片，将调整好的数据拷贝到原来的切片中
	k = k % len(nums)
	res := nums[len(nums)-k:]
	res = append(res, nums[0:len(nums)-k]...)
	copy(nums, res)
}

func sol189_2(nums []int, k int) {
	reverse := func(list []int, start, end int) {
		for start < end {
			list[start], list[end] = list[end], list[start]
			start++
			end--
		}
	}
	// 反转整个数组
	// 反转数组的前k个元素
	// 反转数组的后n-k个元素
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
