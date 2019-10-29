package count

// Count 计数排序
func Count(arr []int) {
	var max int = 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	temps := make([]int, max+1)
	for _, v := range arr {
		temps[v]++
	}
	var id int = 0
	for k, v := range temps {
		for v > 0 {
			arr[id] = k
			id ++
			v--
		}
	}
}
