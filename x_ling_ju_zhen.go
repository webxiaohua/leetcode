/**
 * @author:伯约
 * @date:2023/9/6
 * @note:
零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

示例 1：
输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

示例 2：
输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
**/

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if rows[row] || cols[col] {
				matrix[row][col] = 0
			}
		}
	}
}
