package graph

type Node struct {
	data     interface{}
	children map[int]*Node
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

func (n *Node) GetChildren() map[int]*Node {
	return n.children
}
