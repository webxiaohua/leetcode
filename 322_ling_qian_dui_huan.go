/**
【题目322】零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0


提示：
1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
**/

package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	res := sol322_1(coins, amount)
	fmt.Println(res)
}

func sol322_1(coins []int, amount int) int {
	// f(x) = min(f(x-coin)+1,f(x))
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin { // 当前目标值高于面值
				dp[i] = min(dp[i-coin]+1, dp[i])
				// target = 5 , coins = [2,3]
				//  0 = 0
				//  1 = 0
				//  2 = 1
				//  3 = 1
				//	4 = 2
				//  5 =	2
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
