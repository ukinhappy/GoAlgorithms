package leetcode

func pancakeSort(A []int) []int {

	var result []int
	lenth := len(A)

	for lenth > 0 {
		var index int
		for i, v := range A[:lenth] {
			if v == lenth {
				index = i
			}
		}
		if index != 0 {
			// 先翻转
			trancate(A, 0, index)
			result = append(result, index+1)
		}
		//　再翻转
		trancate(A, 0, lenth-1)
		result = append(result, lenth)
		lenth--
	}
	return result
}

func trancate(A []int, i, j int) {
	for i <= j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}

}
