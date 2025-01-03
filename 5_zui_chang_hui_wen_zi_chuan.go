/**
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
**/

package main

func longestPalindrome(s string) string {
	// 状态转移方程 f(i,j) = f(i+1,j-1) && s[i] == s[j]
	// f(0,0) = true
	// f(0,1) = s[0] == s[1]
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, 0) // true 表示 f(i,j) 是回文串
	// 所有单个字符均为回文串
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	// 先枚举子串长度
	return ""
}
