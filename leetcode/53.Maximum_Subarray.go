package leetcode

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	maxValue := dp[0]
	for i:=1; i<len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		maxValue = max(maxValue, dp[i])
	}
	return maxValue
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}