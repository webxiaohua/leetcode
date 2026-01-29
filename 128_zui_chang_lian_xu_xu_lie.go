/*
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
0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

package main

import "fmt"

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := longestConsecutive(nums)
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	// 题解在于，判断当前元素 -1 是否存在，如果不存在，则表示队首，依次往后找直到没有连续的数字位置，如果存在，则跳出本次循环
	// 引入map，加快速度
	tmpMap := make(map[int]bool)
	for _, num := range nums {
		tmpMap[num] = true
	}
	var maxLength int
	// 取用map 消除重复元素，可以减少循环次数
	for num := range tmpMap {
		if !tmpMap[num-1] {
			currentNum := num
			currentLength := 1
			// 向右扩展序列
			for tmpMap[currentNum+1] {
				currentNum++
				currentLength++
			}
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
