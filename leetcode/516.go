package leetcode

func longestPalindromeSubseq(s string) int {
	return tran(s, 0, len(s)-1)
}

// 递归方式
func tran(s string, i, j int) int {
	if i > j {
		return 0
	}
	if i == j {
		return 1
	}
	if s[i] == s[j] {
		return tran(s, i+1, j-1) + 2
	}
	return max(tran(s, i+1, j), tran(s, i, j-1))
}


//longestPalindromeSubseq dp 方式
func longestPalindromeSubseq(s string) int {
	if s == "" {
		return 0
	}
	var result [][]int
	for i := 0; i < len(s); i++ {
		result = append(result, make([]int, len(s)))
	}
	for i := 0; i < len(s); i++ {
		result[i][i] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				result[i][j] = result[i+1][j-1] + 2
			} else {
				result[i][j] = max(result[i+1][j], result[i][j-1])
			}
		}
	}
	return result[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
