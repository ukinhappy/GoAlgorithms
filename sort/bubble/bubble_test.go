package bubble

import "testing"

func Test_Bubble(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	err := Bubble(arrsort, len(arrsort))
	if err != nil {
		t.Error("bubble failed err:", err)
	}
	t.Log("Bubble 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
	t.Log("BubbleFlage 排序的结果")
	err = BubbleFlage(arrsort, len(arrsort))
	if err != nil {
		t.Error("BubbleFlage failed err:", err)
	}
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
