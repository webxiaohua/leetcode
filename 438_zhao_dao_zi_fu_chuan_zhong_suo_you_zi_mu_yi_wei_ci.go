/**
【题目438】找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

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
	s := "abab" // "cbaebabacd" "abab"
	p := "ab"   // "abc" "ab"
	res := sol438_2(s, p)
	fmt.Println(res)
}

func sol438_1(s string, p string) []int {
	res := make([]int, 0)
	// 双指针每次取 len(p) ,对比是否全在里面
	left := 0
	right := left + len(p) - 1 // 左闭右闭
	pMap := make(map[byte]int) // 字符:数量
	pArr := []byte(p)
	for _, v := range pArr {
		if _, ok := pMap[v]; ok {
			pMap[v]++
		} else {
			pMap[v] = 1
		}
	}
	sArr := []byte(s)
	for right < len(sArr) {
		tmpMap := make(map[byte]int)
		for i := left; i <= right; i++ {
			if _, ok := tmpMap[sArr[i]]; ok {
				tmpMap[sArr[i]]++
			} else {
				tmpMap[sArr[i]] = 1
			}
		}
		// 两个map对比
		allIn := true
		for pV, pCnt := range pMap {
			tmpCnt, tmpOk := tmpMap[pV]
			if !tmpOk {
				allIn = false
				break
			} else if pCnt != tmpCnt {
				allIn = false
				break
			}
		}
		if allIn {
			res = append(res, left)
		}
		left++
		right++
	}
	return res
}

func sol438_2(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	// 双指针每次取 len(p) ,对比是否全在里面
	//left := 0
	//right := left + len(p) - 1 // 左闭右闭
	pMap := make(map[byte]int) // 字符:数量
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}
	// 初始化滑动窗口
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
