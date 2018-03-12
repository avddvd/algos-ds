package graph

import "testing"

func TestNode(t *testing.T) {
	n := NewNode(1)
	if n.GetIntData() != 1 {
		t.Error("error setting/getting data")
	}
	nbr := NewNode(2)
	if nbr.GetData().(int) != 2 {
		t.Error("error setting/getting data")
	}
	n.AddNeighbor(nbr, 10)
	nbrs := n.GetNeighbors()
	if len(nbrs) == 0 {
		t.Error("error adding/getting neighbors")
	}
	for nbr, weight := range nbrs {
		if nbr.GetIntData() != 2 && weight != 10 {
			t.Error("error adding/getting neighbors")
		}
	}
}
