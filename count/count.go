package count

func Count(arr []int, nLength int) {
	if nil == arr || nLength <= 0 {
		return
	}
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
