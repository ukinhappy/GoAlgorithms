package merge

import "testing"

func Test_Merge(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Merge(arrsort, 0, len(arrsort)-1)
	t.Log("Merge 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
