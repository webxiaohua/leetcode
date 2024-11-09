/**
【题目560】和为K的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
**/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	res := sol560_1(nums, k)
	fmt.Println(res)
}

func sol560_1(nums []int, k int) int {
	res := 0
	// 前缀和方式求解
	// 数学论证: 如果存在 k = sum[i-1] + nums[i]； 说明 i 即为求解
	// map[前缀和]次数
	sum := 0
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefixMap[sum-k]; ok {
			res += prefixMap[sum-k]
		}
		prefixMap[sum]++
	}
	return res
}

// 滑动窗口求解
func sol560_2(nums []int, k int) int {

}
