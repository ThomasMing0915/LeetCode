package main

import "fmt"

var dp map[int]int

func coinChange(coins []int, amount int) int {
	dp = make(map[int]int)
	c := coinChangeInner(coins, amount)
	if c < 0x7fffffff {
		return c
	}

	return -1
}

func coinChangeInner(coins []int, amount int) int {
	if count, ok := dp[amount]; ok {
		return count
	}

	if amount < 0 {
		return 0x7fffffff
	}
	if amount == 0 {
		return 0
	}

	minCoins := 0x7fffffff
	for i := 0; i < len(coins); i++ {
		curCoins := coinChangeInner(coins, amount-coins[i]) + 1
		minCoins = min(minCoins, curCoins)
	}

	dp[amount] = minCoins

	return minCoins
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}
