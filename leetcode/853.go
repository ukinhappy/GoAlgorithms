package leetcode

import "sort"

func carFleet(target int, position []int, speed []int) int {

	type tmp struct {
		position int
		tm       float64
	}
	var ps []tmp
	for i, v := range position {
		ps = append(ps, tmp{v, float64((target - position[i])) / float64(speed[i])})
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].position < ps[j].position
	})

	// 根据每个点的剩余时间 求峰顶个数
	var count int = 0
	var max float64 = 0
	for i := len(ps) - 1; i >= 0; i-- {
		if max < ps[i].tm {
			max = ps[i].tm
			count++
		}
	}
	return count
}
