package leetcode

##1.动态规划

var dp = make(map[int]bool)

func canJump(nums []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}
	if index > len(nums)-1 {
		return false
	}
	if dp[index] {
		return true
	}
	var result = false
	for i := 1; i <= nums[i]; i++ {
		result = result || canJump(nums, index+i)
	}
	dp[index] = result
	return dp[index]
}


## 2. 贪心算法
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var last int
	for i := 0; i < len(nums)-1; i++ {
		last = max(nums[i]+i, last)
		if last <= i {
			return false
		}
	}
	return last >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
