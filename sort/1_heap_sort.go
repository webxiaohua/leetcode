/**
 * @author:伯约
 * @date:2025/1/7
 * @note:堆排序
**/

package main

import "fmt"

func maxHeapify(nums []int, i int) {
	if i >= len(nums)/2 {
		return
	}
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < len(nums) {
		if nums[largest] < nums[left] {
			largest = left
		}
	}
	if right < len(nums) {
		if nums[largest] < nums[right] {
			largest = right
		}
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
	}
	i++
	maxHeapify(nums, i)
}

func heapSort(nums []int) {
	maxHeapify(nums, 0)
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums[:i], 0)
	}
}

func main() {
	nums := []int{2, 6, 4, 9, 1, 0, 8}
	heapSort(nums)
	fmt.Println(nums)
}
