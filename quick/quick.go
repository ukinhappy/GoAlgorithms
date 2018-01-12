package quick

func Quick(arr []int, nLow int, nHight int) {
	if arr == nil || nLow < 0 || nHight < 0 {
		return
	}

	if nLow < nHight {
		nMid := sort(arr, nLow, nHight)
		Quick(arr, nLow, nMid-1)
		Quick(arr, nMid+1, nHight)
	}
}

func sort(arr []int, nLow int, nHight int) int {
	var temp int
	temp = arr[nLow]

	// 从后往前找比临时值小的数据
	for nLow < nHight {
		for nLow < nHight {
			if arr[nHight] < temp {
				arr[nLow] = arr[nHight]
				nLow++
				break
			}
			nHight--
		}

		for nLow < nHight {
			if arr[nLow] > temp {
				arr[nHight] = arr[nLow]
				nHight--
				break
			}
			nLow++

		}
	}
	arr[nLow] = temp

	// 从前往后找比临时值大的数据
	return nLow
}

func QuickImprove(arr []int, nLow int, nHight int) {
	if arr == nil || nLow < 0 || nHight < 0 {
		return
	}

	if nLow < nHight {
		nMid := sort(arr, nLow, nHight)
		Quick(arr, nLow, nMid-1)
		Quick(arr, nMid+1, nHight)
	}
}

func sortImprove(arr []int, nLow int, nHight int) int {
	// 将待比较的标准值放到最后边
	if nLow != nHight {
		temp := arr[nLow]
		arr[nLow] = arr[nHight]
		arr[nHight] = temp
	}

	// 记录比标准值小的最后一个元素的下标
	var flag int = nLow - 1
	for ; nLow < nHight; nLow++ {
		if nLow <= nHight {
			flag++
			if flag != nLow {
				arr[nLow] = arr[nLow] ^ arr[flag]
				arr[flag] = arr[nLow] ^ arr[flag]
				arr[nLow] = arr[nLow] ^ arr[flag]
			}
		}
	}
	flag++
	if flag != nHight {
		arr[nHight] = arr[nHight] ^ arr[flag]
		arr[flag] = arr[nHight] ^ arr[flag]
		arr[nHight] = arr[nHight] ^ arr[flag]

	}
	return flag
}
