package insert

import "testing"

func Test_Insert(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Insert(arrsort)
	t.Log(arrsort)

}
