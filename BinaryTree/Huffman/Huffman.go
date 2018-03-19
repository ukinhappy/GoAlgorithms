package main

import (
	"fmt"
)

type Huf struct {
	value  int
	pleft  *Huf
	pright *Huf
}

// 中序遍历
func Print(tree *Huf) {
	if tree == nil {
		return
	}
	Print(tree.pleft)
	fmt.Println(tree.value)
	Print(tree.pright)
}
func createHuf(arr []int) *Huf {
	var root, temp, tempRoot *Huf
	for len(arr) > 0 {
		if root == nil {
			vleft := popmin(arr, len(arr))
			pleft := createHufNode(vleft, nil, nil)
			arr = arr[:len(arr)-1]

			vright := popmin(arr, len(arr))
			pright := createHufNode(vright, nil, nil)
			arr = arr[:len(arr)-1]
			root = createHufNode(vleft+vright, pleft, pright)
		} else {
			min1 := popmin(arr, len(arr))
			arr = arr[:len(arr)-1]
			min2 := topmin(arr, len(arr))
			if min2 <= root.value && len(arr) > 0 {
				min2 = popmin(arr, len(arr))
				arr = arr[:len(arr)-1]

				vleft := popmin(arr, len(arr))
				pleft := createHufNode(min1, nil, nil)
				vright := popmin(arr, len(arr))
				pright := createHufNode(min2, nil, nil)
				temp = createHufNode(vleft+vright, pleft, pright)
				tempRoot = createHufNode(temp.value+root.value, pleft, pright)

			} else {
				temp = createHufNode(min1, nil, nil)
				tempRoot = createHufNode(min1+root.value, nil, nil)
			}

			if root.value > temp.value {
				tempRoot.pleft = temp
				tempRoot.pright = root
			} else {
				tempRoot.pright = temp
				tempRoot.pleft = root
			}
			root = tempRoot
		}

	}

	return root
}

func createHufNode(value int, pleft *Huf, pright *Huf) *Huf {
	return &Huf{value: value, pleft: pleft, pright: pright}
}

func main() {

	arr := []int{4, 9, 23, 435, 7}
	root := createHuf(arr)
	Print(root)
}
