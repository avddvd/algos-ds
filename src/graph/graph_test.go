package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()
	node1 := g.AddEdge(1, 2, 10)
	nodes := g.GetNodes()
	if len(nodes) == 0 {
		t.Error("error adding edge to graph")
	}
	for value, node := range nodes {
		if node.GetIntData() != value {
			t.Error("error adding node/edge to graph")
		}
	}

	visited := map[*Node]bool{}
	DFS(node1, visited)
	if len(visited) == 0 {
		t.Error("error in DFS of a node")
	}
	visited = map[*Node]bool{}
	BFS(node1, visited)
	if len(visited) == 0 {
		t.Error("error in DFS of a node")
	}
}

func TestTopologicalSort(t *testing.T) {
	g := NewGraph()
	g.AddEdge(5, 2, 0)
	g.AddEdge(5, 0, 0)
	g.AddEdge(4, 0, 0)
	g.AddEdge(4, 1, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)
	s := g.TopologicalSort()
	for value, node := range g.GetNodes() {
		for nbr := range node.GetNeighbors() {
			fmt.Printf("%d -> %d \n", value, nbr.GetIntData())
		}
	}
	fmt.Println(s)
	for !s.IsEmpty() {
		n := s.Pop().GetData().(*Node)
		fmt.Println(n.GetIntData())
	}
}
