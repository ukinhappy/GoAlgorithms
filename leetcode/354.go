package leetcode

import "sort"
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}
		if envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
			return true
		}
		return false

	})
	var dp []int
	for i := 0; i < len(envelopes); i++ {
		dp = append(dp, 1)
	}
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	var max = 1
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
