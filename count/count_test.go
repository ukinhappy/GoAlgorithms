package count

import "testing"

func Test_Count(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Count(arrsort, len(arrsort))
	t.Log("Count 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
