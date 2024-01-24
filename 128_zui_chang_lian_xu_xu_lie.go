/**
【题目128】最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
**/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := sol128_2(nums)
	fmt.Println(res)
}

func sol128_1(nums []int) int {
	maxLength := 0
	// 先排序，再循环求解
	sort.Ints(nums)
	tmpLength := 0
	for i, num := range nums {
		if i == 0 {
			tmpLength++
			continue
		}
		if num-nums[i-1] == 0 {
			continue
		} else if num-nums[i-1] == 1 {
			tmpLength++
		} else {
			if maxLength < tmpLength {
				maxLength = tmpLength
			}
			tmpLength = 1
		}
	}
	if maxLength < tmpLength {
		maxLength = tmpLength
	}
	return maxLength
}

func sol128_2(nums []int) int {
	// 构造map  数字:是否出现，遍历map,如果当前数字前面还存在值，则跳过，如果不存在，一直往后找到不连续数字为止
	maxLength := 0
	tmpMap := make(map[int]bool)
	for _, num := range nums {
		tmpMap[num] = true
	}
	for num, _ := range tmpMap {
		if !tmpMap[num-1] {
			currentLength := 1
			tmpNum := num
			for tmpMap[tmpNum+1] {
				currentLength++
				tmpNum++
			}
			if maxLength < currentLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
