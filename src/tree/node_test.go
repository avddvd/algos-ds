package tree

import (
	"testing"
)

func TestNode(t *testing.T) {
	n := new(Node)
	data := 0
	left := NewNode(1)
	right := NewNode(2)
	n.SetData(data)
	if n.GetData().(int) != data {
		t.Error("error setting/getting data for tree node")
	}
	n.SetLeft(left)
	n.SetRight(right)
	if n.GetLeft().GetData().(int) != left.GetData().(int) {
		t.Error("error setting/getting left for tree node")
	}
	if n.GetRight().GetData().(int) != right.GetData().(int) {
		t.Error("error setting/getting right for tree node")
	}
}
