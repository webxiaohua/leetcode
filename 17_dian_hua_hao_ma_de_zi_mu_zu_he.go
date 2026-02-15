/**
题目17:电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
2->abc , 3->def, 4->ghi 5->jkl, 6->mno, 7->pqrs, 8->tuv, 9->wxyz

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = "2"
输出：["a","b","c"]

提示：
1 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

**/

package main

import "fmt"

func main() {
	digits := "23"
	res := sol17_3(digits)
	fmt.Println(res)
}

// 回溯求解 - 深度优先
func sol17_1(digits string) []string {
	// 先简历map，利用数组下表
	phoneMap := []string{
		"", "", // 0和1没有字母
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	var res []string   // 存放结果
	var current []byte // 存放当前组合
	var backtrack func(idx int)
	backtrack = func(idx int) { // idx 指遍历 digits 的索引
		if idx == len(digits) {
			res = append(res, string(current))
			return
		}
		// 获取当前数字对应的字符，如 digits[idx] 是 '2'，则对应下标2
		digit := digits[idx] - '0'
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			current = append(current, letters[i])
			backtrack(idx + 1)
			// 回溯，撤销选择，移除最后一个字母
			current = current[:len(current)-1]
		}
	}
	backtrack(0)
	return res
}

// 循环求解 - 广度优先
func sol17_2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := []string{
		"", "", // 0和1没有字母
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		// 索引下标
		digit := digits[i] - '0'
		// 当前字符串
		letters := phoneMap[digit]
		// 临时字符串
		var tmp []string
		// 遍历已有结果集中的每一个组合
		for _, s := range res {
			// 遍历当前数字的每一个字母，拼接成新组合
			for _, c := range letters {
				tmp = append(tmp, s+string(c))
			}
		}
		// 更新结果集为本轮的新组合
		res = tmp
	}
	return res
}

// 递归求解
func sol17_3(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := []string{
		"", "", // 0和1没有字母
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	// 分治思想 , 把问题拆分成子问题
	splitFn := func(letters string) []string {
		letterArr := []string{}
		for _, letter := range letters {
			letterArr = append(letterArr, string(letter))
		}
		return letterArr
	}
	// 递归的退出条件
	if len(digits) == 1 {
		letters := phoneMap[digits[0]-'0']
		return splitFn(letters)
	}
	var res []string
	frontLetterArr := splitFn(phoneMap[digits[0]-'0'])
	leftLetterArr := sol17_3(digits[1:])
	for _, f := range frontLetterArr {
		for _, b := range leftLetterArr {
			res = append(res, f+b)
		}
	}
	return res
}
