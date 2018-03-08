package bucket

func findMinMax(arr []int) (min, max int) {

	min = arr[0]
	max = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func Bucket(arr []int, nLength int) {

	if arr == nil || nLength <= 0 {
		return
	}
	// 找到最大值和最小值
	min, max := findMinMax(arr)
	bucketlength := max/13 - min/13 + 1

	// 创建
	bucket := make([][]int, bucketlength)
	for i := 0; i < bucketlength; i++ {
		bucket[i] = make([]int, 0)
	}

	// 插入
	for _, v := range arr {

		bucket[v/13-min/13] =
			append(bucket[v/13-min/13], v)
	}

	//对每个进行排序
	for i := 0; i < bucketlength; i++ {
		Quick(bucket[i], 0, len(bucket[i])-1)
	}
	var j int = 0
	for i := 0; i < bucketlength; i++ {
		for _, v := range bucket[i] {
			arr[j] = v
			j++
		}
	}

	return
}

/// 辅助排序
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
