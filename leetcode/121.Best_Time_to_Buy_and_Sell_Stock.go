package leetcode

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]
	for i:=1; i<len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-min)
		if prices[i] < min {
			min = prices[i]
		}
	}

	return dp[len(prices)-1]
}
