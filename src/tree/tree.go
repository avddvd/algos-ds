package tree

import (
	"fmt"
	"github.com/avddvd/algos-ds/queue"
	"github.com/avddvd/algos-ds/stack"
	"math"
)

type Tree struct {
	root *Node
}

func (t *Tree) SetRoot(root *Node) {
	t.root = root
}

func (t *Tree) GetRoot() *Node {
	return t.root
}

func PrintInOrder(root *Node) {
	if root != nil {
		PrintInOrder(root.GetLeft())
		fmt.Println(root.GetData())
		PrintInOrder(root.GetRight())
	}
}

func PrintPreOrder(root *Node) {
	if root != nil {
		fmt.Println(root.GetData())
		PrintPreOrder(root.GetLeft())
		PrintPreOrder(root.GetRight())
	}
}

func PrintPostOrder(root *Node) {
	if root != nil {
		PrintPostOrder(root.GetLeft())
		PrintPostOrder(root.GetRight())
		fmt.Println(root.GetData())
	}
}

func (n *Node) FindMin() *Node {
	for n.GetLeft() != nil {
		n = n.GetLeft()
	}
	return n
}

func (n *Node) FindMax() *Node {
	for n.GetRight() != nil {
		n = n.GetRight()
	}
	return n
}

func (n *Node) Search(data int) bool {
	found := false
	if n.GetData().(int) == data {
		found = true
	} else {
		if n.GetData().(int) > data {
			found = n.GetLeft().Search(data)
		} else {
			found = n.GetRight().Search(data)
		}
	}
	return found
}

func (n *Node) Insert(node *Node) {
	var p *Node
	data := node.GetData().(int)
	for n != nil {
		if n.GetData().(int) < data {
			p = n
			n = n.GetRight()
		} else {
			p = n
			n = n.GetLeft()
		}
	}
	if data < p.GetData().(int) && p.GetLeft() == nil {
		p.SetLeft(node)
	} else if p.GetRight() == nil {
		p.SetRight(node)
	}
}

func IsBSTHelper(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	return node.GetIntData() < max &&
		node.GetIntData() > min &&
		IsBSTHelper(node.GetLeft(), min, node.GetIntData()) &&
		IsBSTHelper(node.GetRight(), node.GetIntData(), max)
}

func (n *Node) IsBST() bool {
	if n == nil {
		return true
	}
	min, max := math.MinInt64, math.MaxInt64
	return IsBSTHelper(n, min, max)
}

func SerializeHelper(node *Node, s []int) []int {
	if node == nil {
		return s
	}
	s = append(s, node.GetIntData())
	s = SerializeHelper(node.GetLeft(), s)
	s = SerializeHelper(node.GetRight(), s)
	return s
}

func (n *Node) Serialize() []int {
	s := []int{}
	return SerializeHelper(n, s)
}

func DeserializeHelper(q *queue.Queue, min, max int) (n *Node) {
	if q.IsEmpty() {
		return
	}
	if q.Peek() > min && q.Peek() < max {
		data := q.Dequeue().GetData().(int)
		n = NewNode(data)
		left := DeserializeHelper(q, min, data)
		if left != nil {
			n.left = left
		}
		right := DeserializeHelper(q, data, max)
		if right != nil {
			n.right = right
		}
	}
	return
}

func Deserialize(list []int) *Node {
	min, max := math.MinInt64, math.MaxInt64
	q := new(queue.Queue)
	for _, item := range list {
		q.Enqueue(item)
	}
	return DeserializeHelper(q, min, max)
}

func FindSuccessor(root, node *Node) *Node {
	var succ *Node
	if node.GetRight() != nil {
		return node.GetRight().FindMin()
	}
	for root != nil {
		if node.GetIntData() < root.GetIntData() {
			succ = root
			root = root.GetLeft()
		} else if node.GetIntData() > root.GetIntData() {
			root = root.GetRight()
		} else {
			break
		}
	}
	return succ
}

func FindPredecessor(root, node *Node) *Node {
	var preDec *Node
	if node.GetLeft() != nil {
		return node.GetLeft().FindMax()
	}
	for root != nil {
		if node.GetIntData() < root.GetIntData() {
			root = root.GetLeft()
		} else if node.GetIntData() > root.GetIntData() {
			preDec = root
			root = root.GetRight()
		} else {
			break
		}
	}
	return preDec
}

// interviewcake problem
// superbalanced : diff b/n any leaf node should be less than <= 1
