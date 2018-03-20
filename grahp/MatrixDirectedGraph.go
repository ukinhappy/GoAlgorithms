/*
邻接矩阵有向图Graph( Matrix Directed Graph)：采用数组的方式存储
为了方便处理，我们假设最大有a-z 26个顶点
使用我们提供的数据建图

DFS 深度优先便利
BFS 广度优先遍历
DFSByStack 通过栈 深度便利

相对于无向图的DFS，在无向图的基础上，将未遍历到的基础上，继续将未遍历到的点进行遍历。BFS同DFS
*/

package main

import (
	"fmt"
	"github.com/ukinhappy/go-tools/queue"
	"github.com/xcltapestry/xclpkg/algorithm"
)

type Graph struct {
	nVertex int //顶点
	nEdge   int //边
	Matrix  [][]int
	Vexs    []rune
}

func main() {

	var vexs []rune = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	var edges [][2]rune = [][2]rune{{'A', 'C'}, {'A', 'D'}, {'A', 'F'}, {'B', 'C'}, {'C', 'D'}, {'E', 'G'}, {'F', 'G'}}
	graph := createGrahp(vexs, edges)

	fmt.Println("  A B C D E F G")
	for i := 0; i < len(graph.Matrix); i++ {
		fmt.Printf("%c", 'A'+i)
		fmt.Println(graph.Matrix[i])
	}

	BFS(graph)
	fmt.Printf("\n")
	DFS(graph)
	fmt.Printf("\n")
	DFSByStack(graph)
}

func createGrahp(vexs []rune, edges [][2]rune) *Graph {
	var graph *Graph = &Graph{}
	graph.Vexs = make([]rune, len(vexs))
	graph.Matrix = make([][]int, len(vexs))
	for i := 0; i < len(vexs); i++ {
		graph.Matrix[i] = make([]int, len(vexs))
	}
	graph.nVertex = len(vexs)
	graph.nEdge = len(edges)
	for i := 0; i < len(edges); i++ {
		x := edges[i][0] - 'A'
		y := edges[i][1] - 'A'
		graph.Matrix[x][y] = 1
		//  graph.Matrix[y][x] = 1
	}
	return graph
}

func BFS(graph *Graph) {
	if graph == nil {
		return
	}

	flage := make(map[int]bool, graph.nVertex)
	for i := 0; i < graph.nVertex; i++ {
		flage[i] = false
	}
	// 设定A为开始点，也可以不设定
	bfs(graph, flage, 0)
	for flageindex, f := range flage {
		if !f {
			bfs(graph, flage, flageindex)
		}

	}
}

func bfs(graph *Graph, flage map[int]bool, flageindex int) {
	queue := queue.New()
	queue.Push(flageindex)
	for !queue.Empty() {
		index := queue.Pop().(int)
		if !flage[index] {
			fmt.Printf("%c", 'A'+index)
			flage[index] = true
			for k, v := range graph.Matrix[index] {
				if v == 1 && !flage[k] {
					queue.Push(k)
				}
			}
		}
	}
}

func DFS(graph *Graph) {
	if graph == nil {
		return
	}
	flage := make(map[int]bool, graph.nVertex)
	for i := 0; i < graph.nVertex; i++ {
		flage[i] = false
	}
	// 设定A为开始点，也可以不设定
	dfs(graph, 0, flage)

	for flageindex, f := range flage {
		if !f {
			dfs(graph, flageindex, flage)
		}
	}
}

func dfs(graph *Graph, num int, flage map[int]bool) {
	if graph == nil {
		return
	}

	if !flage[num] {
		flage[num] = true
		for k, v := range graph.Matrix[num] {
			if !flage[k] && v == 1 {
				dfs(graph, k, flage)
			}
		}
		fmt.Printf("%c", 'A'+num)
	}

}

func DFSByStack(graph *Graph) {
	flage := make(map[int]bool, graph.nVertex)

	stack := algorithm.NewStack()
	stack.Push(0)
	flage[0] = true
	p := true
	for !stack.Empty() {
		p = true
		index := stack.Top().(int)

		for k, v := range graph.Matrix[index] {
			if !flage[k] && v == 1 {
				flage[k] = true
				p = false
				stack.Push(k)
				break
			}
		}

		if p {
			fmt.Printf("%c", 'A'+index)
			stack.Pop()
		}
	}
}
