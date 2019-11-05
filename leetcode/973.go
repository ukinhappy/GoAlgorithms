package leetcode

import "sort"

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return dist(points[i]) < dist(points[j])
	})

	return points[:K]
}
func dist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
