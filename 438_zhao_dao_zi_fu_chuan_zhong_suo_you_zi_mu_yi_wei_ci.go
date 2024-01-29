/**
【题目438】找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词:指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
**/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := sol438_2(s, p)
	fmt.Println(res)
}

func sol438_1(s string, p string) []int {
	// 滑动窗口，双重循环，复杂度太高
	res := make([]int, 0)
	start, end := 0, len(p)-1
	arr := []byte(s)
	targetMap := sol438_map([]byte(p))
	for end < len(s) {
		currentMap := sol438_map(arr[start : end+1])
		if reflect.DeepEqual(currentMap, targetMap) {
			res = append(res, start)
		}
		start++
		end++
	}
	return res
}

func sol438_map(arr []byte) (res map[byte]int) {
	res = make(map[byte]int)
	for i := 0; i < len(arr); i++ {
		res[arr[i]]++
	}
	return res
}

func sol438_2(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	pMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}
	window := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		window[s[i]]++
	}
	if reflect.DeepEqual(window, pMap) {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		if window[s[i-len(p)]] > 1 {
			window[s[i-len(p)]]--
		} else {
			delete(window, s[i-len(p)])
		}
		window[s[i]]++
		if reflect.DeepEqual(window, pMap) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
