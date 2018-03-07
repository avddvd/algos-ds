package tree

import (
	"fmt"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	n := NewNode(3)
	for _, val := range []int{1, 2, 4, -1, 0} {
		node := NewNode(val)
		n.Insert(node)
	}
	n.Insert(NewNode(-1))
	PrintInOrder(n)
}

func TestIsBST(t *testing.T) {
	n := NewNode(10)
	for _, val := range []int{1, 2, 4, -1, 0} {
		node := NewNode(val)
		n.Insert(node)
	}
	PrintInOrder(n)
	if n.IsBST() != true {
		t.Error("error in isBST method")
	}

	nt := NewNode(5)
	nt.SetRight(NewNode(-1))
	nt.SetLeft(NewNode(10))
	if nt.IsBST() != false {
		t.Error("error in isBST method")
	}
}

func TestSerilaize(t *testing.T) {
	n := NewNode(10)
	for _, val := range []int{1, 2, 4, -1, 0} {
		node := NewNode(val)
		n.Insert(node)
	}
	PrintInOrder(n)
	if n.IsBST() != true {
		t.Error("error in isBST method")
	}
	s := n.Serialize()
	fmt.Println(s)
	if len(n.Serialize()) == 0 {
		t.Error("error serializing a bst")
	}
	r := Deserialize(s)
	PrintInOrder(r)
	if r == nil {
		t.Error("error deserializing tree")
	}
}

func TestSuccessorAndPredec(t *testing.T) {
	n := NewNode(10)
	for _, val := range []int{1, 2, 4, -1, 0} {
		node := NewNode(val)
		n.Insert(node)
	}
	n1 := NewNode(15)
	n2 := NewNode(-15)
	n.Insert(n1)
	n.Insert(n2)
	PrintInOrder(n)
	s := FindSuccessor(n, n2)
	p := FindPredecessor(n, n1)
	fmt.Println(s.GetIntData(), p.GetIntData())

}
