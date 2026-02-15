/**
26. 删除有序数组中的重复项

给你一个 非严格递增排列 的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的相对顺序应该保持一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
nums 的前 k 个元素应包含排序后的唯一数字。下标 k - 1 之后的剩余元素可以忽略。

判题标准:
系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4,_,_,_,_,_]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

提示：
1 <= nums.length <= 3 * 10^4
-100 <= nums[i] <= 100
nums 已按非递减顺序排列。
**/

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := sol26_2(nums)
	fmt.Println(res)
	fmt.Println(nums)
}

func sol26_1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 双指针
	i, j := 0, 1
	for {
		if j >= len(nums) {
			break
		}
		if nums[i] > nums[j] {
			break
		}
		if nums[i] == nums[j] {
			// 性能比较差，数组需要来回移动
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i++
			j++
		}
	}
	return j
}

func sol26_2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
