package main

func topmin(arr []int, nLength int) int {
	if len(arr) <= 0 || nLength <= 0 {
		return -1
	}
	if nLength == 1 {
		return arr[0]
	}
	// 将原有的数组调整成小堆的样子
	// 调整可以从根节点往下调整，也可以从最后一个非叶子节点往上调整
	for i := nLength/2 - 1; i >= 0; i-- {
		adjust(arr, i, nLength)
	}
	return arr[0]
}

func popmin(arr []int, nLength int) int {
	if len(arr) <= 0 || nLength <= 0 {
		return -1
	}
	if nLength == 1 {
		return arr[0]
	}
	// 将原有的数组调整成小堆的样子
	// 调整可以从根节点往下调整，也可以从最后一个非叶子节点往上调整
	for i := nLength/2 - 1; i >= 0; i-- {
		adjust(arr, i, nLength)
	}
	arr[0] = arr[0] ^ arr[nLength-1]
	arr[nLength-1] = arr[0] ^ arr[nLength-1]
	arr[0] = arr[0] ^ arr[nLength-1]
	return arr[nLength-1]
}
func adjust(arr []int, adjustIndex int, nLength int) {
	for {
		// 当前节点有两个孩子
		if 2*adjustIndex+2 < nLength {
			// 右节点小于左节点
			if arr[2*adjustIndex+2] < arr[2*adjustIndex+1] {
				// 右节点与父节点交换
				if arr[2*adjustIndex+2] > arr[adjustIndex] {
					arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+2]
					arr[2*adjustIndex+2] = arr[adjustIndex] ^ arr[2*adjustIndex+2]
					arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+2]
					// 调整右节点为待调整的树
					adjustIndex = 2*adjustIndex + 2
				} else {
					break
				}
			} else {
				if arr[2*adjustIndex+1] < arr[adjustIndex] {
					// 左节点与父节点交换
					arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
					arr[2*adjustIndex+1] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
					arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
					// 调整左节点为待调整的树
					adjustIndex = 2*adjustIndex + 1
				} else {
					break
				}
			}
		} else if 2*adjustIndex+1 < nLength { //  当前节点有一个节点（左节点）
			// 子节点小于父节点
			if arr[2*adjustIndex+1] < arr[adjustIndex] {
				arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
				arr[2*adjustIndex+1] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
				arr[adjustIndex] = arr[adjustIndex] ^ arr[2*adjustIndex+1]
				// 调整左节点为待调整的树
				adjustIndex = 2*adjustIndex + 1
			} else {
				break
			}
		} else {
			// 没有子节点 无需调整
			break
		}
	}
}
