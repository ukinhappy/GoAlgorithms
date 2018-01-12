package bubble

import (
	"errors"
	//	"fmt"
)

func Bubble(arr []int, nLength int) error {
	if nil == arr {
		return errors.New("slice is null")
	}

	for i := 0; i < nLength-1; i++ {
		for j := 0; j < nLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return nil
}

func BubbleFlage(arr []int, nLength int) error {
	if nil == arr {
		return errors.New("slice is null")
	}

	for i := 0; i < nLength-1; i++ {
		var flage int = 0
		for j := 0; j < nLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				flage = j
				arr[j] = arr[j] ^ arr[j+1]
				arr[j+1] = arr[j] ^ arr[j+1]
				arr[j] = arr[j] ^ arr[j+1]
			}
		}
		if 0 == flage {
			return nil
		}
		i = nLength - 1 - flage - 1
	}

	return nil
}
