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
	temp := arr[nLow]

	for nLow < nHight {
		//　从后往前找第一个比temp　小的值
		for nLow < nHight {
			if arr[nHight] < temp {
				arr[nLow] = arr[nHight]
				nLow ++
				break
			}
			nHight --
		}
		//　从前往后找第一个比temp　大的值
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
	return nLow
}

func QuickImprove(arr []int, nLow int, nHight int) {
	if arr == nil || nLow < 0 || nHight < 0 {
		return
	}

	if nLow < nHight {
		nMid := sortImprove(arr, nLow, nHight)
		Quick(arr, nLow, nMid-1)
		Quick(arr, nMid+1, nHight)
	}
}

func sortImprove(arr []int, nLow int, nHight int) int {
	var flag = nLow
	arr[nLow], arr[nHight] = arr[nHight], arr[nLow]

	for nLow < nHight {
		if arr[nLow]<arr[nHight]{
			if flag!=nLow{
				arr[flag],arr[nLow]=arr[nLow],arr[flag]
			}
			flag ++
		}
		nLow++
	}
	if flag!=nHight{
		arr[flag],arr[nHight]=arr[nHight],arr[flag]
	}
	return flag
}
