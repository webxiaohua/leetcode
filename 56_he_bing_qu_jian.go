/**
【题目56】合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4
**/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	res := sol56_1(a)
	fmt.Println(res)
}

func sol56_1(intervals [][]int) [][]int {
	res := make([][]int, 0)
	// 先按 starti 进行升序排序，然后创建结果数组，遍历原数组，每次跟结果数组的最后一条记录进行对比，判断是否重合
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, item := range intervals {
		if len(res) == 0 {
			res = append(res, item)
			continue
		}
		if res[len(res)-1][1] < item[0] {
			res = append(res, item)
		} else if res[len(res)-1][1] > item[1] {
			continue
		} else {
			res[len(res)-1][1] = item[1]
		}
	}
	return res
}
