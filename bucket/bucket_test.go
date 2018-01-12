package bucket

import "testing"

func Test_Bucket(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Bucket(arrsort, len(arrsort))

	t.Log("Bucket 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
