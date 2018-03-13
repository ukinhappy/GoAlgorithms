package main

import (
	"fmt"
)

type Node struct {
	pLeft  *Node
	pRight *Node
	hight  int
	value  int
}

type AVL struct {
	tree *Node
}

func main() {
	var arr = []int{5, 6, 23, 64, 31, 33}

	var avl AVL
	avl.CreateTree(arr)

	Print(avl.tree)
	DeleteNode(avl.tree, 5)
	fmt.Println("------------------------------")
	Print(avl.tree)
	DeleteNode(avl.tree, 64)
	fmt.Println("------------------------------")
	Print(avl.tree)
}

func (avl *AVL) CreateTree(arr []int) {
	for _, v := range arr {
		avl.tree = InsertNode(avl.tree, v)
	}
}

func InsertNode(tree *Node, num int) *Node {
	if tree == nil {
		tree = &Node{value: num}
	}

	//插入到左子树
	if num < tree.value {
		tree.pLeft = InsertNode(tree.pLeft, num)
		if getHight(tree.pLeft)-getHight(tree.pRight) >= 2 {
			if num > tree.pLeft.value {
				// ll
				tree = LL(tree)
			} else {
				// lR
				tree = LR(tree)
			}

		}

	} else if num > tree.value {
		tree.pRight = InsertNode(tree.pRight, num)
		if getHight(tree.pRight)-getHight(tree.pLeft) >= 2 {
			if num > tree.pRight.value {
				// RR
				tree = RR(tree)
			} else {
				// Rl
				tree = RL(tree)
			}
		}
	}

	tree.hight = max(getHight(tree.pLeft), getHight(tree.pRight)) + 1
	return tree
}

func LL(tree *Node) *Node {
	var n1 *Node = tree
	var n2 *Node = tree.pLeft
	n1.pLeft = n2.pRight
	n2.pRight = n1
	n1.hight = max(getHight(n1.pLeft), getHight(n1.pRight))
	n2.hight = max(getHight(n2.pLeft), getHight(n2.pRight))
	return n2
}
func LR(tree *Node) *Node {
	tree.pLeft = RR(tree.pLeft)
	return LL(tree)
}
func RR(tree *Node) *Node {
	var n1 *Node = tree
	var n2 *Node = tree.pRight
	n1.pRight = n2.pLeft
	n2.pLeft = n1
	n1.hight = max(getHight(n1.pLeft), getHight(n1.pRight))
	n2.hight = max(getHight(n2.pLeft), getHight(n2.pRight))
	return n2
}
func RL(tree *Node) *Node {
	tree.pRight = LL(tree.pRight)
	return RR(tree)
}

func getHight(tree *Node) int {
	if tree != nil {
		return tree.hight
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 中序遍历
func Print(tree *Node) {
	if tree == nil {
		return
	}
	Print(tree.pLeft)
	fmt.Println(tree.value)
	Print(tree.pRight)
}

func DeleteNode(tree *Node, num int) *Node {
	if tree == nil {
		return nil
	}

	if num < tree.value {
		// 在左边
		tree.pLeft = DeleteNode(tree.pLeft, num)
		if getHight(tree.pRight)-getHight(tree.pLeft) >= 2 {
			if getHight(tree.pRight.pLeft) > getHight(tree.pRight.pRight) {
				// RL
				tree = RL(tree)
			} else {
				tree = RR(tree)
			}
		}
	} else if num > tree.value {
		tree.pRight = DeleteNode(tree.pRight, num)
		if getHight(tree.pLeft)-getHight(tree.pRight) >= 2 {
			if getHight(tree.pLeft.pLeft) > getHight(tree.pLeft.pRight) {
				tree = LL(tree)
			} else {
				tree = LR(tree)
			}
		}
	} else {
		// 要删除的元素
		if tree.pLeft != nil && tree.pRight != nil {
			// 找左子树和右子树最高的
			if getHight(tree.pLeft) > getHight(tree.pRight) {
				// 找左子树的最大值
				max := findMax(tree.pLeft)
				tree.value = max
				DeleteNode(tree.pLeft, max)
			} else {
				min := findMin(tree.pRight)
				tree.value = min
				DeleteNode(tree.pRight, min)
			}

		} else {

			if tree.pLeft != nil {
				tree = tree.pLeft
			} else {
				tree = tree.pRight
			}
		}
	}
	return tree
}

func findMax(tree *Node) int {
	if tree == nil {
		return -1
	}
	for tree.pRight != nil {
		tree = tree.pRight
	}
	return tree.value
}

func findMin(tree *Node) int {
	if tree == nil {
		return -1
	}
	for tree.pLeft != nil {
		tree = tree.pLeft
	}
	return tree.value
}
