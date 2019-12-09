package leetcode

import "sort"

func reorganizeString(S string) string {
	ss := []byte(S)
	if len(ss) <= 1 {
		return S
	}

	length := len(ss)
	ssort := make(map[byte]int, 0)
	var maxlength int = 0
	for _, v := range ss {
		ssort[v] ++
		if ssort[v] > maxlength {
			maxlength = ssort[v]
		}
	}
	if maxlength > (length+1)/2 {
		return ""
	}
	sort.Slice(ss, func(i, j int) bool {
		if ssort[ss[i]] > ssort[ss[j]] {
			return true
		} else if ssort[ss[i]] == ssort[ss[j]] {
			return ss[i] > ss[j]
		}
		return false
	})

	result := make([]byte, length)
	var index int = 0
	for _, v := range ss {
		if index >= length {
			index = 1
		}
		result[index] = v
		index += 2
		v--
	}

	return string(result)
}
