/**
 4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
**/

package main

import "fmt"

func main() {
	/*nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))*/
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortedArr := make([]int, 0)
	if len(nums1) == 0 {
		sortedArr = append(sortedArr, nums2...)
	} else if len(nums2) == 0 {
		sortedArr = append(sortedArr, nums1...)
	} else {
		i, j := 0, 0
		for i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				sortedArr = append(sortedArr, nums1[i])
				i++
			} else {
				sortedArr = append(sortedArr, nums2[j])
				j++
			}
		}
		if i < len(nums1) {
			sortedArr = append(sortedArr, nums1[i:]...)
		}
		if j < len(nums2) {
			sortedArr = append(sortedArr, nums2[j:]...)
		}
	}
	if len(sortedArr) == 0 {
		return 0
	}
	if len(sortedArr)%2 == 0 {
		return float64(sortedArr[len(sortedArr)/2]+sortedArr[len(sortedArr)/2-1]) / 2
	} else {
		return float64(sortedArr[len(sortedArr)/2])
	}
}
