package graph

import "github.com/avddvd/algos-ds/queue"

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
	var node *Node
	if _, ok := g.nodes[data]; ok {
		node = g.nodes[data]
	}
	return node
}

func (g *Graph) AddEdge(from, to *Node, weight int) {
	if g.GetNode(from.GetIntData()) == nil {
		g.AddNode(from)
	}

	if g.GetNode(to.GetIntData()) == nil {
		g.AddNode(to)
	}
	from.AddNeighbor(to, weight)
}

func dfs(node *Node, visited map[*Node]bool) {
	if _, ok := visited[node]; ok {
		return
	}
	visited[node] = true
	for nbr := range node.GetNeighbors() {
		dfs(nbr, visited)
	}
}

func bfs(node *Node, visited map[*Node]bool) {
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
