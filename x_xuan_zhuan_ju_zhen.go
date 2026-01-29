/**
 * @author:伯约
 * @date:2023/9/3
 * @note:
【旋转矩阵】
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

**/

package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(arr)
	for _, item := range arr {
		for _, item2 := range item {
			fmt.Print(item2, " ")
		}
		fmt.Println("")
	}
	fmt.Println(arr)
}

func rotate(arr [][]int) {
	// 分析过程，四个点位变化 5*5  row=1 col=2
	// a[row][col] -> a[col][n-row-1]  a[0,1] = a[1,2]
	// a[col][n-row-1] -> a[n-row-1][n-col-1]  a[2,3] = a[3,2]
	// a[n-row-1][n-col-1] -> a[col][n-row-1]  a[3,2] = a[2,1]
	// a[col][n-row-1] -> a[row][col]
	n := len(arr)
	for row := 0; row < n/2; row++ {
		for col := 0; col < (n+1)/2; col++ {
			tmp := arr[row][col]
			arr[row][col] = arr[n-col-1][row]
			arr[n-col-1][row] = arr[n-row-1][n-col-1]
			arr[n-row-1][n-col-1] = arr[col][n-row-1]
			arr[col][n-row-1] = tmp
		}
	}
}
