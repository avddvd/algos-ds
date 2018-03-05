package tree

import "fmt"

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
