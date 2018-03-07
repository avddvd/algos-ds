package node

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) SetPrev(prev *Node) {
	n.prev = prev
}

func (n *Node) GetPrev(prev *Node) *Node {
	return n.prev
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
