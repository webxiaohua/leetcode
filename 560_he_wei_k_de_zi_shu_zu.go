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
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7
**/

package main

import "fmt"

func main() {
	/*nums := []int{1, 1, 1}
	k := 2*/
	/*nums := []int{1, 2, 3}
	k := 3*/
	/*nums := []int{1, 2, 1, 2, 1}
	k := 3*/
	nums := []int{1, -1, 0}
	k := 0
	res := sol560_2(nums, k)
	fmt.Println(res)
}

func sol560_1(nums []int, k int) int {
	res := 0
	// 前缀和map求解 前缀和:数量， 遍历数组，求当前元素的前缀和，如果能找到 本次前缀和 - k 的元素存在于map，则给结果追加数量
	// 由于数据自身可能已经等于目标值，所以前缀和初始化后需要设置前缀和为0的key，值为1
	// 数学论证： 第i个元素的前缀和 = 前i-1的元素和 + 第i个元素
	// s[i] = s[i-1] + a[i];
	// 如果存在 s[i] - k  == s[i-1] , 则 a[i] = k , 属于有效解
	sum := 0                          // 累计前缀和
	prefixSumMap := make(map[int]int) // {"前缀和":数量}
	prefixSumMap[0] = 1               // 表示前缀和为 0 的情况出现 1 次，处理边界情况，比如数组第一个元素就等于 k（此时 pre_sum = k，pre_sum -k =0，哈希表中存在 0，计数 + 1）。
	for _, num := range nums {
		sum += num
		// 如果存在前缀和为 sum-k ，说明从前面某一位置到当前位置和为 k
		if cnt, ok := prefixSumMap[sum-k]; ok {
			res += cnt
		}
		prefixSumMap[sum]++
	}
	return res
}

/*
arr: 1,2,3,4,
sum: 1,3,6,10

arr[i,j] = sum[j] - sum[i-1] , i>0

a[1,2] = 5
sum[2] - sum[0] = 5

map{sum_val:cnt}
{
0:1
1:1
3:1
6:1
10:1
}
k=5


sum[prefix] - k ?

arr[i,j] = sum[j] - sum[i-1]
sum[j] - sum[i-1] = k => sum[i-1] = sum[j] - k
翻译： 如果存在和为k的子数组 arr[i,j]，相当于 sum[j]-k 刚好等于 sum[i-1]
用map来存的话，上面的式子可表现为： sumMap[curSumVal - k]++
*/

func sol560_2(nums []int, k int) int {
	// [1,-1,0] 0
	res := 0
	// 构建 sumMap
	sumMap := make(map[int]int) // {sumVal:cnt}
	sumMap[0] = 1               // 兜底值, 和为0表示1次，如果数组第一个值刚好等于k
	tmpSum := 0
	for _, num := range nums {
		tmpSum += num
		if _, ok := sumMap[tmpSum-k]; ok {
			res += sumMap[tmpSum-k]
		}
		sumMap[tmpSum]++
	}
	return res
}
