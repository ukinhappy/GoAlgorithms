package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// start排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] {
			tmp[1] = max(tmp[1], intervals[i][1])
			continue
		}
		result = append(result, tmp)
		tmp = intervals[i]
	}
	return append(result, tmp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
