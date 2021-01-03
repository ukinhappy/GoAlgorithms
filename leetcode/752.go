package leetcode

func openLock(deadends []string, target string) int {
	if target=="0000"{
		return 0
	}
	deadendsM := make(map[string]bool, 0)
	for _, v := range deadends {
		deadendsM[v] = true
	}
	if deadendsM[target] {
		return -1
	}
	if deadendsM["0000"] {
		return -1
	}
	depth := 1
	queue := []string{"0000"}
	deadendsM["0000"] = true
	for {
		tmpQueue := make([]string, 0)
		if len(queue) <= 0 {
			break
		}
		for _, v := range queue {
			// 获取v 可达到的所有值
			for i := 0; i < 4; i++ {
				tmp := up(v, i)
				if tmp == target {
					return depth
				}
				if !deadendsM[tmp] {
					tmpQueue = append(tmpQueue, tmp)
					deadendsM[tmp]=true
				}
				tmp = down(v, i)
				if tmp == target {
					return depth
				}
				if !deadendsM[tmp] {
					tmpQueue = append(tmpQueue, tmp)
					deadendsM[tmp]=true
				}
			}
		}

		queue = tmpQueue
		depth++
	}

	return -1
}

func up(v string, n int) string {
	res := []rune(v)
	if res[n] == '9' {
		res[n] = '0'
	} else {
		res[n] += 1
	}
	return string(res)
}

func down(v string, n int) string {
	res := []rune(v)
	if res[n] == '0' {
		res[n] = '9'
	} else {
		res[n] -= 1
	}
	return string(res)
}
