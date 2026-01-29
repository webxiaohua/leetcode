/**
【题目49】字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]

提示：
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

**/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	tmpMap := make(map[string][]string)
	// 排序 + map[string][]string
	for _, str := range strs {
		strArr := strings.Split(str, "")
		sort.Strings(strArr)
		sortedStr := strings.Join(strArr, "")
		if _, ok := tmpMap[sortedStr]; ok {
			tmpMap[sortedStr] = append(tmpMap[sortedStr], str)
		} else {
			tmpMap[sortedStr] = []string{str}
		}
	}
	for _, list := range tmpMap {
		res = append(res, list)
	}
	return res
}
