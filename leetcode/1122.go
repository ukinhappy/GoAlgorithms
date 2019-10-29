package leetcode

func relativeSortArray(arr1 []int, arr2 []int) []int {

	var max = 0
	for _, v := range arr1 {
		if max < v {
			max = v
		}
	}
	tmps := make([]int, max+1)
	for _, v := range arr1 {
		tmps[v]++
	}

	var result []int
	for _, v := range arr2 {
		for tmps[v] > 0 {
			result = append(result, v)
			tmps[v]--
		}
	}
	for i, v := range tmps {
		for v > 0 {
			result = append(result, i)
			v--
		}
	}
	return result
}
