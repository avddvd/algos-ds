package node

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) GetNext() *Node {
	var next *Node
	if n != nil {
		next = n.next
	}
	return next
}

func (n *Node) GetData() interface{} {
	return n.data
}
