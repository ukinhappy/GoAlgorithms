package count

// Count 计数排序
func Count(arr []int) {
	nLength := len(arr)
	var max int = 0
	for i := 0; i < nLength; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	var arrTemp []int = make([]int, max+1)

	for i := 0; i < nLength; i++ {
		arrTemp[arr[i]]++
	}
	var j int = 0
	for i := 0; i < max+1; i++ {
		for arrTemp[i] > 0 {
			arr[j] = i
			arrTemp[i]--
			j++
		}

	}
}
