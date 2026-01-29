/*
【题目0】旋转字符串

给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部，
如把字符串“abcdef”前面的2个字符’a’和’b’移动到字符串的尾部，使得原字符串变成字符串“cdefab”。
请写一个函数完成此功能，要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)。

*/

package main

import "fmt"

func main() {
	// step1 先翻转整个字符串  12345 -> 54321
	// step2 翻转前 n-k 个字符 54321 -> 34521
	// step3 翻转后k个字符 34521 -> 34512
	source := "12345"
	res := reverseString([]byte(source), 0, len(source)-1)
	res = reverseString(res, 0, 2)
	res = reverseString(res, 3, 4)
	fmt.Println(string(res))
	fmt.Println("Success")
}

// 翻转字符串
func reverseString(s []byte, start, end int) []byte {
	if start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
		reverseString(s, start, end)
	}
	return s
}
