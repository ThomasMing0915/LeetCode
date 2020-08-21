package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	//dp[i]表示以nums[i]结尾的最大前缀和
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	max := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
