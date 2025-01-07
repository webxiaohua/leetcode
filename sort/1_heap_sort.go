/**
 * @author:伯约
 * @date:2025/1/7
 * @note:堆排序
**/

package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 3, 5, 9}
	heapSort(nums)
	fmt.Println(nums)
}

// 堆化处理 (nums 数组, n表示当前需要处理的数组长度, i表示当前元素)
func maxHeapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, n, largest)
	}
}

// 构建大顶堆
func buildMaxHeap(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(nums, n, i)
	}
}

// 堆排序
func heapSort(nums []int) {
	// 构建大顶堆
	buildMaxHeap(nums)
	for i := 0; i < len(nums); i++ {
		// 将堆顶元素与末尾元素交换
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		// 重新构建大顶堆
		maxHeapify(nums, len(nums)-1-i, 0)
	}
}
