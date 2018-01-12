package quick

import "testing"

func Test_Quick(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Quick(arrsort, 0, len(arrsort)-1)
	t.Log("Quick 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}

func Test_QuickImprove(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	QuickImprove(arrsort, 0, len(arrsort)-1)
	t.Log("QuickImprove 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
