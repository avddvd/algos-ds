package graph

type Node struct {
	data      interface{}
	neighbors map[*Node]int
}

func NewNode(data interface{}) *Node {
	return &Node{
		data:      data,
		neighbors: map[*Node]int{},
	}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) GetData() interface{} {
	return n.data
}

func (n *Node) GetIntData() int {
	return n.data.(int)
}

func (n *Node) GetNeighbors() map[*Node]int {
	return n.neighbors
}

func (n *Node) AddNeighbor(nbr *Node, weight int) {
	n.neighbors[nbr] = weight
}
