package heap

import "testing"

func Test_Heap(t *testing.T) {
	arrsort := []int{4, 9, 23, 435, 7}
	Heap(arrsort, len(arrsort))
	t.Log("Heap 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
