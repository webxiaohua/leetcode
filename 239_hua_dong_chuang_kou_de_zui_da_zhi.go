/**
【题目239】滑动窗口的最大值
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

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := sol239_1(nums, k)
	fmt.Println(res)
}

func sol239_1(nums []int, k int) []int {
	res := make([]int, 0)
	// 双端队列 队首存放最大值的索引，维护当前滑动窗口的有效数据索引
	deque := make([]int, 0)
	pushQueue := func(i int) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	// 初始化双端队列
	for i := 0; i < k; i++ {
		pushQueue(i)
	}
	// 开始滑动
	for i := k; i < len(nums); i++ {
		res = append(res, nums[deque[0]])
		// 移除滑动窗口外的元素
		for len(deque) > 0 && i-deque[0] >= k {
			deque = deque[1:]
		}
		pushQueue(i)
	}
	res = append(res, nums[deque[0]])
	return res
}
