package bubble

import "testing"

func Test_Bubble(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	Bubble(arrsort)
	t.Log(arrsort)


}

func Test_BubbleFlage(t *testing.T) {
	arrsort := []int{2, 45, 1, 22, 4, 56, 77, 82, 84, 85, 21}
	BubbleFlage(arrsort)
	t.Log(arrsort)
}
