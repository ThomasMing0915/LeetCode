package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	//初始化初始值
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][len(dp[i])-1] = dp[i-1][len(dp[i-1])-1] + triangle[i][len(triangle[i])-1]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}

	minValue := dp[len(dp)-1][0]
	for i := 1; i < len(dp[len(dp)-1]); i++ {
		if minValue > dp[len(dp)-1][i] {
			minValue = dp[len(dp)-1][i]
		}
	}
	//fmt.Println(dp)
	return minValue
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {
	c := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(c))
}
