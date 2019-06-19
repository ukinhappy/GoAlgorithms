package count

import "testing"

func Test_Count(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Count(arrsort)

	t.Log(arrsort)
}
