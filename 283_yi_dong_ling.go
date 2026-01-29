/*
【题目283】移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

提示:
1 <= nums.length <= 104
-2^31 <= nums[i] <= 2^31 - 1

进阶：你能尽量减少完成的操作次数吗？
*/
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{1, 0}
	sol_283_2(nums)
	fmt.Println(nums)
}

// 双指针，快指针从第一个元素开始，往后查找，直到找到第一个不为0的数字，
// 随后慢指针开始往后找，找到第一个为0的数字，两者调换位置，继续从快指针指向的下标往后
func sol_283_1(nums []int) {
	var zeroPointer, nonZeroPointer int
	for {
		if nonZeroPointer == len(nums) || zeroPointer == len(nums) {
			return
		}
		if nums[nonZeroPointer] != 0 {
			if nums[zeroPointer] == 0 {
				// 调换顺序
				if nonZeroPointer > zeroPointer {
					nums[nonZeroPointer], nums[zeroPointer] = nums[zeroPointer], nums[nonZeroPointer]
				}
				nonZeroPointer = zeroPointer + 1
			} else {
				zeroPointer++
			}
		} else {
			nonZeroPointer++
		}
	}
}

// 指针i用于迭代，指针 nonZeroIndex 用于记录当前非0元素的指针，循环一遍后，所有非0元素已经归位，nonZeroIndex 之后的元素再置为0
func sol_283_2(nums []int) {
	var nonZeroIndex int
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}
	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
