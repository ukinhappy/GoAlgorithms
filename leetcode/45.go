package leetcode

import (
	"fmt"
	"math"
)

## 1. 动态规划
var dp = make(map[int]int)

func canJump(nums []int, index int) int {
	if index >= len(nums)-1 {
		return 0
	}
	if dp[index]!=0{
		return dp[index]
	}
	var res int = math.MaxInt64
	for i := 1; i <= nums[index]; i++ {
		res = min(canJump(nums, index+i)+1, res)
	}
	dp[index]=res
	return dp[index]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
