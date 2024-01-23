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

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口（遍历，比较） + 双指针（计算长度）
	// 定义map {元素:最后一次出现位置}
	maxLength := 0
	mapLastIndex := make(map[byte]int)
	// 双指针，定义开头指针
	start := 0
	// 遍历，开启滑动窗口
	for end, val := range []byte(s) {
		// 如果在map中出现过，说明已经存在，滑动窗口左侧指针往右移动到上一次出现的位置
		// 同时需要考虑 abba 的问题， 必需当前下标超过start指针才行
		if lastIndex, ok := mapLastIndex[val]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		// 记录元素上一次出现的位置
		mapLastIndex[val] = end
		// 跟maxLength比较
		if (end - start + 1) > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
