package leetcode

import "sort"

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) == len(d[j]) {
			return d[i] < d[j]
		}
		return len(d[i]) > len(d[j])
	})
	for _, v := range d {
		j := 0
		for i := 0; i < len(s); i++ {
			if s[i] == v[j] {
				j++
			}
			if len(v) == j {
				return v
			}
		}
	}
	return ""
}
