/*
go 实现二叉搜素树的添加节点，删除节点，中序遍历
*/
package main

import (
	"fmt"
)

type Tree struct {
	Value  int
	Pleft  *Tree
	Pright *Tree
}

type BST struct {
	tree *Tree
}

// 创建二叉搜索树
func (bst *BST) CreateBinarySearchTree(arr []int) {
	for _, v := range arr {
		bst.InsertNode(v)
	}
}

func (bst *BST) InsertNode(num int) {
	var tree *Tree = bst.tree
	if tree == nil {
		bst.tree = &Tree{Value: num}
		return
	}
	for {
		if tree.Value < num {
			if tree.Pright == nil {
				tree.Pright = &Tree{Value: num}
				return
			}
			tree = tree.Pright
		} else {
			if tree.Pleft == nil {
				tree.Pleft = &Tree{Value: num}
				return
			}
			tree = tree.Pleft
		}
	}
}

// 删除节点
func (bst *BST) FindDelNode(num int) (*Tree, *Tree) {
	if bst.tree == nil {
		return nil, nil
	}
	var pFather, pDel *Tree
	pDel = bst.tree
	for pDel != nil {
		if pDel.Value == num {
			return pDel, pFather
		}
		pFather = pDel
		if pDel.Value > num {
			pDel = pDel.Pleft
		} else {
			pDel = pDel.Pright
		}
	}
	return pDel, pFather
}

func (bst *BST) DeleteNode(num int) {
	// 查找需要删除的节点以及父节点
	pDel, pDelFather := bst.FindDelNode(num)
	if pDel == nil {
		return
	}

	// 两个子节点 找右子树的最大值
	if pDel.Pleft != nil && pDel.Pright != nil {
		// 记录下当前这个点
		bj := pDel
		pDelFather = pDel
		pDel = pDel.Pright
		for pDel.Pleft != nil {
			pDelFather = pDel
			pDel = pDel.Pleft
		}
		bj.Value = pDel.Value

		if pDelFather.Pleft == pDel {
			pDelFather.Pleft = pDel.Pright
		} else {
			pDelFather.Pright = pDel.Pright
		}

	} else { //一个子节点或者没有子节点直接删除
		// 根节点
		if pDelFather == nil {
			if pDel.Pright != nil {
				bst.tree = pDel.Pright
			} else {
				bst.tree = pDel.Pleft
			}
		}
		if pDelFather.Pleft == pDel {
			if pDel.Pright != nil {
				pDelFather.Pleft = pDel.Pright
			} else {
				pDelFather.Pleft = pDel.Pleft
			}
		} else {
			if pDel.Pright != nil {
				pDelFather.Pright = pDel.Pright
			} else {
				pDelFather.Pright = pDel.Pleft
			}
		}
	}
}

type List struct {
	Pstart *Tree
	Pend   *Tree
}

// 排序二叉树转变成双向链表
func (list *List) TreeToList(tree *Tree) {
	if tree == nil {
		return
	}

	list.TreeToList(tree.Pleft)

	if list.Pstart == nil {
		list.Pstart = tree
	} else {
		list.Pend.Pright = tree
		tree.Pleft = list.Pend
	}
	list.Pend = tree
	list.TreeToList(tree.Pright)
}

// 中序遍历
func Print(tree *Tree) {
	if tree == nil {
		return
	}
	Print(tree.Pleft)
	fmt.Println(tree.Value)
	Print(tree.Pright)
}

func main() {
	var arr = []int{6, 3, 78, 2, 5, 66}
	var bst BST
	bst.CreateBinarySearchTree(arr)
	Print(bst.tree)
	fmt.Println("--------------删除节点------------------")
	bst.DeleteNode(5)
	Print(bst.tree)
	fmt.Println("--------------树变链表------------------")
	var list List
	list.TreeToList(bst.tree)
	for list.Pstart != nil {
		fmt.Println(list.Pstart.Value)
		list.Pstart = list.Pstart.Pright
	}
}
