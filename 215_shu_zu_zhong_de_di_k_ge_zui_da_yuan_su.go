/*
【题目215】数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	res := findKthLargest(nums, k) // 5
	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
	res := 0
	// 初始化小根堆
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	// 遍历数组，将元素加入小根堆
	for _, num := range nums {
		heap.Push(minHeap, num)
		// 当小根堆的元素个数大于 k 时，弹出堆顶元素
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	// 返回堆顶元素
	res = (*minHeap)[0]
	return res
}

// 定义小根堆
type MinHeap []int

// 升序排列
func (m MinHeap) Len() int {
	return len(m)
}
func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 实现 heap.Interface 接口
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() any {
	old := *m
	top := old[len(old)-1]
	*m = old[:len(old)-1]
	return top
}
