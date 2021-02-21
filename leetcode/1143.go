package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	var result [][]int
	for i := 0; i <= len(text1); i++ {
		var tmp []int
		for j := 0; j <= len(text2); j++ {
			tmp = append(tmp, 0)
		}
		result = append(result, tmp)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				result[i][j] = result[i-1][j-1] + 1
			} else {
				result[i][j] = max(result[i-1][j], result[i][j-1])
			}
		}

	}
	return result[len(text1)][len(text2)]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
