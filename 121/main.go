package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]

	dp := make([]int, len(prices))

	for i := 1; i < len(prices); i++ {
		diffPrice := prices[i] - minPrice
		if diffPrice > 0 {
			dp[i] = diffPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	maxPrice := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > maxPrice {
			maxPrice = dp[i]
		}
	}
	return maxPrice
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := 0x7fffffff
	maxprofit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxprofit {
				maxprofit = profit
			}
		}
	}
	return maxprofit
}

func main() {
	tests := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}

	for _, test := range tests {
		fmt.Println(maxProfit2(test))
	}
}
