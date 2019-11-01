package leetcode

func isAnagram(s string, t string) bool {

	ss := []byte(s)
	tt := []byte(t)
	if len(ss) != len(tt) {
		return false
	}

	tmp := make(map[byte]int)
	for i := 0; i < len(ss); i++ {
		tmp[ss[i]] = tmp[ss[i]] + 1
		tmp[tt[i]] = tmp[tt[i]] - 1
		if tmp[ss[i]] == 0 {
			delete(tmp, ss[i])
		}
		if tmp[tt[i]] == 0 {
			delete(tmp, tt[i])
		}
	}
	if len(tmp) <= 0 {
		return true
	}
	return false
}
