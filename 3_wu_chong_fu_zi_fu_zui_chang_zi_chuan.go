/**
 3. 无重复字符的最长子串
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
	s := "dvdf"
	res := sol3_1(s)
	fmt.Println(res)
}

func sol3_1(s string) int {
	res, start, end := 0, 0, 0
	tmpMap := make(map[byte]int)
	tmpArr := []byte(s)
	for ; end < len(s); end++ {
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
