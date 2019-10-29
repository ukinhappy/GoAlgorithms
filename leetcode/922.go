package leetcode

func sortArrayByParityII(A []int) []int {
	var o, j = 0, 1
	for o = 0; o < len(A); o += 2 {
		if A[o]%2 == 1 {
			for i := j; i < len(A); i += 2 {
				if A[i]%2 == 0 {
					A[o], A[i] = A[i], A[o]
					j = i
					break
				}
			}
		}
	}

	return A
}
