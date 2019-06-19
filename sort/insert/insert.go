/*Package insert .....
	稳定排序
	时间复杂度:O(N2)
 */
package insert

//Insert 插入排序
func Insert(arr []int) {

	// 从第二个元素开始往前比较
	for i := 1; i < len(arr); i++ {
		lastID := i - 1
		tmp := arr[i]
		for lastID >= 0 && arr[lastID] > tmp {
			arr[lastID+1] = arr[lastID]
			lastID --
		}
		arr[lastID+1] = tmp
	}

}
