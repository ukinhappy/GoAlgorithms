package leetcode

//## 1. 递归
var back map[int]int

func rob(nums []int) int {
	back = make(map[int]int)
	return dp(nums, 0)
}
func dp(nums []int, index int) int {
	if index > len(nums)-1 {
		return 0
	}
	if back[index] != 0 {
		return back[index]
	}
	back[index] = max(nums[index]+dp(nums, index+2), dp(nums, index+1))
	return back[index]
}

//# 2. dp
func rob2(nums []int) int {
	dp := make([]int, len(nums)+2)

	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}
	return dp[0]
}

//# 3. dp 优化
func rob3(nums []int) int {
	dp2 := 0
	dp1 := 0
	dp := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp = max(nums[i]+dp2, dp1)
		dp2 = dp1
		dp1 = dp
	}
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
