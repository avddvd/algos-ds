package graph

import (
	"github.com/avddvd/algos-ds/queue"
	"github.com/avddvd/algos-ds/stack"
)

type Graph struct {
	nodes map[int]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: map[int]*Node{},
	}
}

func (g *Graph) GetNodes() map[int]*Node {
	return g.nodes
}

func (g *Graph) AddNode(node *Node) {
	g.nodes[node.GetIntData()] = node
}

func (g *Graph) GetNode(data int) *Node {
	return g.nodes[data]
}

func (g *Graph) AddEdge(from, to int, weight int) *Node {
	var n1, n2 *Node
	if _, ok := g.nodes[from]; !ok {
		n1 = NewNode(from)
		g.AddNode(n1)
	} else {
		n1 = g.nodes[from]
	}

	if _, ok := g.nodes[to]; !ok {
		n2 = NewNode(to)
		g.AddNode(n2)
	} else {
		n2 = g.nodes[to]
	}
	n1.AddNeighbor(n2, weight)
	return n1
}

func DFS(node *Node, visited map[*Node]bool) {
	if _, ok := visited[node]; ok {
		return
	}
	visited[node] = true
	for nbr := range node.GetNeighbors() {
		DFS(nbr, visited)
	}
}

func BFS(node *Node, visited map[*Node]bool) {
	q := new(queue.Queue)
	q.Enqueue(node)
	for !q.IsEmpty() {
		node := q.Dequeue().GetData().(*Node)
		if _, ok := visited[node]; ok {
			break
		}
		visited[node] = true
		for nbr := range node.GetNeighbors() {
			q.Enqueue(nbr)
		}
	}
}

func topologicalSortHelper(node *Node,
	visited map[*Node]bool, stk *stack.Stack) {
	if value, ok := visited[node]; ok && value == true {
		return
	}
	visited[node] = true
	for nbr := range node.GetNeighbors() {
		topologicalSortHelper(nbr, visited, stk)
	}
	stk.Push(node)
}

func (g *Graph) TopologicalSort() *stack.Stack {
	visited := map[*Node]bool{}
	stk := new(stack.Stack)
	for _, node := range g.GetNodes() {
		if value, ok := visited[node]; !(ok && value == true) {
			topologicalSortHelper(node, visited, stk)
		}
	}
	return stk
}
