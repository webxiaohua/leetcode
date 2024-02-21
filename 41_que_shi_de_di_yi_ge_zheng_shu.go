/**
【题目41】缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1

提示：
1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1
**/

package main

import "fmt"

func main() {
	nums := []int{3, 8, -1, 1}
	res := sol41_2(nums)
	fmt.Println(res)
}

func sol41_1(nums []int) int {
	res := 1
	tmpMap := make(map[int]int)
	for _, num := range nums {
		tmpMap[num] = 0
	}
	for i := 0; i < len(tmpMap); i++ {
		if _, ok := tmpMap[res]; !ok {
			break
		}
		res++
	}
	return res
}

func sol41_2(nums []int) int {
	n := len(nums)
	// 将所有小于等于 0 的数都置为数组的长度加一
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 将每个数对应的索引处的数标记为负数（表示该数存在）
	for i := range nums {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 找出第一个正数的索引，该索引加一即为所求的最小未出现的正整数
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	// 如果没有找到正数，则返回数组长度加一
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
