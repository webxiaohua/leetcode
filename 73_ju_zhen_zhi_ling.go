/**
【题目73】矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

提示：
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1

进阶：
一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
**/

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	sol73_1(matrix)
	fmt.Println(matrix)
}

func sol73_1(matrix [][]int) {
	// 遍历二维数组，找到0的 （i,j）
	// 坐标项去重
	// 渲染0
	xserizes := make([]int, 0)
	yserizes := make([]int, 0)
	ylen := len(matrix)
	xlen := len(matrix[0])
	for i, item1 := range matrix {
		for j, item2 := range item1 {
			if item2 == 0 {
				yserizes = append(yserizes, i)
				xserizes = append(xserizes, j)
			}
		}
	}
	for _, x := range xserizes {
		for i := 0; i < ylen; i++ {
			matrix[i][x] = 0
		}
	}
	for _, y := range yserizes {
		for i := 0; i < xlen; i++ {
			matrix[y][i] = 0
		}
	}
}
