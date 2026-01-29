/**
【题目239】滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
**/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := sol239_2(nums, k)
	fmt.Println(res)
}

func sol239_1(nums []int, k int) []int {
	res := make([]int, 0)
	// 求数组最大值，可按照对比方式进行， O(N) 复杂度
	// 滑动窗口相当于在原数组内切出了很多个子数组，需要求解子数组，复杂度 O(N^2)
	if k == 1 {
		return nums
	}
	left, right := 0, k-1
	for right < len(nums) {
		a, b, max := left, right, nums[left]
		for a <= b {
			if nums[a] > max {
				max = nums[a]
			}
			a++
		}
		res = append(res, max)
		left++
		right++
	}
	return res
}

func sol239_2(nums []int, k int) []int {
	res := make([]int, 0)
	// 双端队列，用于快速得出窗口最大值, 存放数组下标
	deque := make([]int, 0)
	push := func(index int) {
		// 确保双端队列的队首为滑动窗口的最大值
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[index] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, index)
	}
	// 初始化第一个滑动窗口
	for i := 0; i < k; i++ {
		push(i)
	}
	for i := k; i < len(nums); i++ {
		res = append(res, nums[deque[0]])
		// 双端队列移除元素
		if len(deque) > 0 && i-deque[0] >= k {
			deque = deque[1:]
		}
		push(i)
	}
	res = append(res, nums[deque[0]])
	return res
}
