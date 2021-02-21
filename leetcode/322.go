package leetcode

func coinChange(coins []int, amount int) int {
	var dp []int
	for i := 0; i <= amount; i++ {
		dp = append(dp, amount+1)
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, v := range coins {
			if i < v {
				continue
			}
			if dp[i-v]+1 < dp[i] {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
