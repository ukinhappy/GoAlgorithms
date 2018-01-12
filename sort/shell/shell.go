package shell

import (
	"fmt"
)

func Shell(arr []int, nLength int) {
	if nil == arr || nLength <= 0 {
		return
	}
	fmt.Println("nLength:", nLength)
	// 确定步长
	for gap := nLength / 2; gap >= 1; gap /= 2 {
		fmt.Println("gap:", gap)
		// 对每一个步长对应的数组进行插入排序
		for i := gap; i < nLength; i += gap {
			flag := i - gap
			temp := arr[i]
			for flag >= 0 && arr[flag] > temp {
				arr[flag+gap] = arr[flag]
				flag -= gap
			}
			arr[flag+gap] = temp
		}
	}
}
