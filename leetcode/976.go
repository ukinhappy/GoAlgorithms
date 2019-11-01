package leetcode

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	var max int = 0
	for i := len(A) - 3; i >= 0; i-- {
		if (A[i]+A[i+1] > A[i+2]) {
			tmp := A[i] + A[i+1] + A[i+2]
			if max < tmp {
				max = tmp
			}
		}
	}

	return max
}
