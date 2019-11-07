package leetcode

import "sort"

func hIndex(ciations []int) int {
	if len(ciations) <= 0 {
		return 0
	}
	sort.Ints(ciations)

	length := len(ciations)
	for i, v := range ciations {
		if length-i <= v {
			return length - i
		}
	}
	return 0
}
