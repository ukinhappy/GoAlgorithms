/*Package bubble ....
	时间复杂度 O(N2)
	稳定排序
 */
package bubble

//Bubble bubble
func Bubble(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}

}

//BubbleFlage bubble flag
func BubbleFlage(arr []int) {
	//
	var changeFlag = len(arr) - 1
	var index = changeFlag
	for changeFlag > 0 {
		changeFlag = 0
		for i := 0; i < index; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				changeFlag = i

			}
		}

		index = changeFlag
	}

}
