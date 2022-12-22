package main

import "fmt"

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	nums = quickSort(nums)
	fmt.Println(nums)
}

// 快速排序
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left, right := 0, len(nums)-1
	isLeft := false
	base := nums[0]
	for {
		if left == right {
			break
		}
		if isLeft {
			if nums[left] > base {
				nums[right] = nums[left]
				right--
				isLeft = false
			} else {
				left++
			}
		} else {
			if nums[right] < base {
				nums[left] = nums[right]
				left++
				isLeft = true
			} else {
				right--
			}
		}
	}
	nums[left] = base
	quickSort(nums[0:left])
	quickSort(nums[left+1:])
	return nums
}

// 归并排序
