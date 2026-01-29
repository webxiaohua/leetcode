/**
387. 字符串中的第一个唯一字符
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1 。

示例 1：
输入: s = "leetcode"
输出: 0

示例 2：
输入: s = "loveleetcode"
输出: 2

示例 3：
输入: s = "aabb"
输出: -1

提示：
1 <= s.length <= 10^5
**/

package main

import "fmt"

func main() {
	fmt.Println(sol387_2("leetcode"))     // 0
	fmt.Println(sol387_2("loveleetcode")) // 2
	fmt.Println(sol387_2("aabb"))         // -1
}

// map法
func sol387_1(s string) int {
	tmpMap := make(map[rune]int)
	for _, char := range s {
		tmpMap[char]++
	}
	for i, char := range s {
		if tmpMap[char] == 1 {
			return i
		}
	}
	return -1
}

// 数组法，利用英文字母仅有26个字符特性
func sol387_2(s string) int {
	arr := [26]uint{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
