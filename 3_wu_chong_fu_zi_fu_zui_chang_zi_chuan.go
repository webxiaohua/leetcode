/**
【题目3】无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

示例 4:
输入: s = "abba"
输出: 2

请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
**/

package main

import "fmt"

func main() {
	//s := "abba"
	//s := "bbbbb"
	//s := "pwwkew"
	s := "abba"
	res := sol3_2(s)
	fmt.Println(res)
}

func sol3_1(s string) int {
	// 滑动窗口
	res, start, end := 0, 0, 0
	tmpMap := make(map[byte]int)
	tmpArr := []byte(s)
	for ; end < len(s); end++ {
		// 需要注意，start 可能跑到相同元素值后面，此时不需要计算
		if _, ok := tmpMap[tmpArr[end]]; ok && start <= tmpMap[tmpArr[end]] {
			res = sol3GetMax(res, end-start)
			start = tmpMap[tmpArr[end]] + 1
		}
		tmpMap[tmpArr[end]] = end
	}
	res = sol3GetMax(res, end-start)
	return res
}

func sol3GetMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func sol3_2(s string) int {
	left, res := 0, 0
	tmpMap := make(map[byte]int)
	tmpArr := []byte(s)
	for right := 0; right < len(tmpArr); right++ {
		if _, ok := tmpMap[tmpArr[right]]; ok && tmpMap[tmpArr[right]] >= left {
			left = tmpMap[tmpArr[right]] + 1
		}
		tmpMap[tmpArr[right]] = right
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
