package leetcode

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	back(n, nil)
	return res
}

func back(n int, track []string) {
	if len(track) == n {
		tmp := make([]string, n)
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		canUse := true
		for j, v := range track {
			if v[i] == 'Q' {
				canUse = false
				break
			}

			var right = i + len(track) - j
			if right < n && v[right] == 'Q' {
				canUse = false
				break
			}
			var left = i - len(track) + j
			if left >= 0 && v[left] == 'Q' {
				canUse = false
				break
			}
		}
		if !canUse {
			continue
		}
		track = append(track, makeString(n, i))
		back(n, track)
		track = track[:len(track)-1]
	}
}
func makeString(n int, index int) string {
	var res []byte
	for i := 0; i < n; i++ {
		res = append(res, '.')
	}
	res[index] = 'Q'
	return string(res)
}

