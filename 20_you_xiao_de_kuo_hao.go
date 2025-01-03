/*
【题目20】有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	每个右括号都有一个对应的相同类型的左括号。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([])"
输出：true

提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

package main

func isValid(s string) bool {
	// 遇到左括号就入栈，遇到右括号出栈并做配对，如果配对失败直接返回false，否则遍历到最后，栈为空则返回true，不空返回false
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			outChar := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if v == ')' && outChar != '(' {
				return false
			} else if v == ']' && outChar != '[' {
				return false
			} else if v == '}' && outChar != '{' {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
