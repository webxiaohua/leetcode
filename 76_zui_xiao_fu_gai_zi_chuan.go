/**
【题目76】最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。


提示：
m == s.length
n == t.length
1 <= m, n <= 10^5
s 和 t 由英文字母组成

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
**/

package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := sol76_1(s, t)
	fmt.Println(res)
}

func sol76_1(s string, t string) string {
	res := ""
	// 如何判断 s 包含 t, 用两个map,key 表示字符，val表示出现次数，当 tMap 完全在 sMap 中时表示存在
	// 求最短可通过滑动窗口，left,right
	maxLength := 99999999
	minLen := 99999999 // 最小长度
	minLenStart := 0   // 最小长度开始索引
	enoughCnt := 0     // 满足条件的字符数量
	tMap := make(map[byte]int)
	currentMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	left, right := 0, 0
	for right < len(s) {
		currentMap[s[right]]++
		if count, exist := tMap[s[right]]; exist && currentMap[s[right]] == count {
			enoughCnt++
		}
		// 判断是否完全满足
		for enoughCnt == len(tMap) {
			// 更新最小子串信息
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLenStart = left
			}
			// 左指针向右移动，缩小窗口
			currentMap[s[left]]--
			if count, exist := tMap[s[left]]; exist && currentMap[s[left]] < count {
				enoughCnt--
			}
			left++
		}
		// 向右移动指针
		right++
	}
	if minLen != maxLength {
		res = s[minLenStart : minLenStart+minLen]
	}
	return res
}
