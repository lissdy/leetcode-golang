package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := i
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && min >= dp[i-coins[j]] {
				min = dp[i-coins[j]]
			}
		}
		dp[i] = min + 1
	}
	fmt.Print(dp)
	return dp[amount]
}

func main() {
	coins := []int{2}
	res := coinChange(coins, 3)
	fmt.Print(res)
}
