package merge

func Merge(arr []int, nLow, nHight int) {
	if nil == arr || nLow < 0 || nHight < 0 {
		return
	}
	if nLow < nHight {
		nMid := (nLow + nHight) / 2
		Merge(arr, nLow, nMid)
		Merge(arr, nMid+1, nHight)

		// 递归到最底层 只剩两个元素
		sort(arr, nLow, nHight)
	}
}

func sort(arr []int, nLow int, nHight int) {
	if nil == arr || nLow < 0 || nHight < 0 {
		return
	}
	// 申请一个临时数组
	var arrtemp []int = make([]int, nHight-nLow+1)

	var nMid int = (nLow + nHight) / 2
	var begin1 int = nLow
	var end1 int = nMid
	var begin2 int = nMid + 1
	var end2 int = nHight
	var i int = 0
	for begin1 <= end1 && begin2 <= end2 {
		// 找两个数组的最小值
		if arr[begin1] < arr[begin2] {
			arrtemp[i] = arr[begin1]
			begin1++
		} else {
			arrtemp[i] = arr[begin2]
			begin2++
		}
		i++
	}

	// 将剩下的某个还没有取完数据的数组内的数据添加到 arrtemp中去
	for ; begin1 <= end1; begin1++ {
		arrtemp[i] = arr[begin1]
		i++
	}

	for ; begin2 <= end2; begin2++ {
		arrtemp[i] = arr[begin2]
		i++
	}

	// 在将临时数组的元素重写写入原数组
	for k, v := range arrtemp {
		arr[nLow+k] = v
	}
}
