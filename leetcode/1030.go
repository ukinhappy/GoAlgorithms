package leetcode

import "fmt"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	l1 := r0 + c0
	l2 := R - r0 - 1 + C - c0 - 1
	l3 := R - r0 - 1 + c0
	l4 := r0 + C - c0 - 1
	var max = l1
	if l2 > max {
		max = l2
	}
	if l3 > max {
		max = l3
	}
	if l4 > max {
		max = l4
	}
	max+=1
	tmp := make([][][]int, max)
	for i := 0; i < max; i++ {
		tmp[i] = make([][]int, 0)
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			l := abs(r0, i) + abs(c0, j)
			fmt.Println(int(l))
			tmp[int(l)] = append(tmp[int(l)], []int{i, j})
		}
	}
	var result [][]int
	for _, v := range tmp {
		if len(v) > 0 {
			result = append(result, v...)
		}
	}
	return result
}

func abs(r, c int) int {
	if r-c < 0 {
		return c - r
	}
	return r - c
}
