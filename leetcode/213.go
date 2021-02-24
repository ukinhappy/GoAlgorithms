package leetcode

func rob1(nums []int) int {
	fmt.Println(nums)
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
func rob(nums []int) int {
	if len(nums)<1{
		return 0
	}
	if len(nums)<2{
		return nums[0]
	}
	return max(rob1(nums[:len(nums)-1]), rob1(nums[1:]))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

