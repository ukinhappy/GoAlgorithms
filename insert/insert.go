package insert

func Insert(arr []int, nLength int) {
	if arr == nil || nLength < 0 {
		return
	}
	var temp, j int
	for i := 1; i < nLength; i++ {
		// 记录已经排好序的最后一个元素的下标
		j = i - 1
		// 记住需要排序的第一个元素
		temp = arr[i]

		// 注意j >=0 在前面
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
