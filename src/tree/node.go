package tree

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) SetLeft(left *Node) {
	n.left = left
}

func (n *Node) SetRight(right *Node) {
	n.right = right
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) GetData() interface{} {
	return n.data
}

func (n *Node) GetIntData() int {
	return n.data.(int)
}
