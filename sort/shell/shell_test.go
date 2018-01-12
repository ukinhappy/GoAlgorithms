package shell

import "testing"

func Test_Shell(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Shell(arrsort, len(arrsort))
	t.Log("Shell 排序的结果")
	for index, value := range arrsort {
		t.Log("K:", index, " value:", value)
	}
}
