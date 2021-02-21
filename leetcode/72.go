package leetcode

func minDistance(word1 string, word2 string) int {
	return tran1(word1, word2, len(word1)-1, len(word2)-1)
}

//tran1 基础方法1 从后往前所有可行方式递归
func tran1(word1, word2 string, i, j int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}

	if word1[i] == word2[j] {
		return tran1(word1, word2, i-1, j-1)
	}

	return min(tran1(word1, word2, i-1, j)+1,
		tran1(word1, word2, i, j-1)+1,
		tran1(word1, word2, i-1, j-1)+1,
	)
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

// DP优化后的方式
func minDistance2(word1 string, word2 string) int {
	var result [][]int
	for i := 0; i <= len(word1); i++ {
		result = append(result, make([]int, len(word2)+1))
	}

	for k, _ := range word2 {
		result[0][k+1] = k + 1
	}

	for k, _ := range word1 {
		result[k+1][0] = k + 1
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				result[i][j] = result[i-1][j-1]
				continue
			}
			result[i][j] = min(result[i-1][j-1]+1, result[i][j-1]+1, result[i-1][j]+1)
		}
	}
	return result[len(word1)][len(word2)]
}
