package leetcode

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	var dp []int

	for i := 0; i < len(nums); i++ {
		dp = append(dp, 0)
	}

	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxFunc(nums[i], dp[i-1]+nums[i])
		max = maxFunc(max, dp[i])
	}
	return max

}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
