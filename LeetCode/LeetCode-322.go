package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	dp := []int{}
	for i := 0; i < amount+1; i++ {
		dp = append(dp, amount+1)
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	res := coinChange(coins, 11)
	fmt.Println(res)
}
