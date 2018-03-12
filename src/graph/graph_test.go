package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()
	node1 := NewNode(1)
	node2 := NewNode(2)
	g.AddEdge(node1, node2, 10)
	nodes := g.GetNodes()
	nbrs := node1.GetNeighbors()
	if len(nodes) == 0 {
		t.Error("error adding edge to graph")
	}
	for value, node := range nodes {
		if node.GetIntData() != value {
			t.Error("error adding node/edge to graph")
		}
	}
	if len(nbrs) == 0 {
		t.Error("error adding neighbors/edges")
	}

	visited := map[*Node]bool{}
	dfs(node1, visited)
	if len(visited) == 0 {
		t.Error("error in dfs of a node")
	}
	fmt.Println(visited)
	visited = map[*Node]bool{}
	bfs(node1, visited)
	if len(visited) == 0 {
		t.Error("error in dfs of a node")
	}
	fmt.Println(visited)
}
