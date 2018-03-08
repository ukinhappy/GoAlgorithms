/*
Author : happy
堆排序:利用堆的特性，将数组调整成大堆的结构，
然后依次将堆顶的元素 从数组的末尾往前放，每拿走一个堆顶的元素就要重新调整堆，保证剩下的元素始终是一个大堆
*/
package heap

func Heap(arr []int, nLength int) {
	if len(arr) <= 0 || nLength <= 0 {
		return
	}
	// 将原有的数组调整成大堆的样子
	// 调整可以从根节点往下调整，也可以从最后一个非叶子节点往上调整
	for i := nLength/2 - 1; i >= 0; i-- {
		adjust(arr, i, nLength)
	}

	// 依次将大堆的堆顶的元素 从放到数组的末尾往前放，并调整剩下的元素，保证是一个大堆
	for i := nLength - 1; i >= 0; i-- {
		arr[0] = arr[0] ^ arr[i]
		arr[i] = arr[0] ^ arr[i]
		arr[0] = arr[0] ^ arr[i]
		// 调整除了最后i个元素，剩下的元素，使其成为大堆
		adjust(arr, 0, i)
	}

}

func adjust(arr []int, adjustIndex int, nLength int) {

	for {
		// 当前节点有两个孩子
		if 2*adjustIndex+2 < nLength {
			// 右节点大于左节点
			if arr[2*adjustIndex+2] > arr[2*adjustIndex+1] {
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
				if arr[2*adjustIndex+1] > arr[adjustIndex] {
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
			// 子节点大于父节点
			if arr[2*adjustIndex+1] > arr[adjustIndex] {
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
