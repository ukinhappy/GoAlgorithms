package leetcode

import (
	"sort"
	"fmt"
)

func rearrangeBarcodes(A []int) []int {
	counts := make([]int, 10001)

	for _, v := range A {
		counts[v]++
	}

	sort.Slice(A, func(i, j int) bool {
		if counts[A[i]] == counts[A[j]] {
			return A[i] > A[j]
		}
		return counts[A[i]] > counts[A[j]]
	})
	fmt.Println(A)
	res := make([]int, len(A))
	var i = 0
	for _, v := range A {
		res[i] = v
		i += 2
		if i >= len(A) {
			i = 1
		}
	}

	return res
}
