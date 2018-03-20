/*
邻接矩阵无向图Graph( list Undirected Graph)：采用数组的方式存储
为了方便处理，我们假设最大有a-z 26个顶点
使用我们提供的数据建图
*/

package main

import (
	"fmt"
)

type Graph struct {
	nVertex  int //顶点
	nEdge    int //边
	vertexs  []rune
	Vexslist []*firstNode
}
type firstNode struct {
	nVertex      rune //顶点
	nVertexslice []rune
}

func main() {

	var vexs []rune = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	var edges [][2]rune = [][2]rune{{'A', 'C'}, {'A', 'D'}, {'A', 'F'}, {'B', 'C'}, {'C', 'D'}, {'E', 'G'}, {'F', 'G'}}
	graph := createGrahp(vexs, edges)

	for _, v := range graph.Vexslist {

		fmt.Printf("%c", v.nVertex)
		for _, vv := range v.nVertexslice {
			fmt.Printf("%c", vv)
		}
		fmt.Printf("\n")
	}

}

func createGrahp(vexs []rune, edges [][2]rune) *Graph {
	var graph *Graph = &Graph{}
	graph.vertexs = make([]rune, len(vexs))
	graph.Vexslist = make([]*firstNode, len(vexs))

	graph.nVertex = len(vexs)
	graph.nEdge = len(edges)

	for i := 0; i < len(edges); i++ {
		x := edges[i][0] - 'A'
		y := edges[i][1] - 'A'

		if graph.Vexslist[x] == nil {
			graph.Vexslist[x] = &firstNode{nVertex: edges[i][0], nVertexslice: make([]rune, 0)}
		}
		graph.Vexslist[x].nVertexslice = append(graph.Vexslist[x].nVertexslice, edges[i][1])

		if graph.Vexslist[y] == nil {
			graph.Vexslist[y] = &firstNode{nVertex: edges[i][1], nVertexslice: make([]rune, 0)}
		}
		graph.Vexslist[y].nVertexslice = append(graph.Vexslist[y].nVertexslice, edges[i][0])
	}
	return graph
}
