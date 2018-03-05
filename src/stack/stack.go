package stack

import (
	"github.com/avddvd/algos-ds/node"
)

type Stack struct {
	top *node.Node
}

func (s *Stack) Push(data interface{}) {
	node := new(node.Node)
	node.SetData(data)
	if s.top == nil {
		s.top = node
	} else {
		node.SetNext(s.top)
		s.top = node
	}
}

func (s *Stack) Pop() *node.Node {
	var top *node.Node
	if s.top != nil {
		top = s.top
		s.top = top.GetNext()
	}
	return top
}

func (s *Stack) Peek() interface{} {
	var top interface{}
	if s.top != nil {
		top = s.top.GetData()
	}
	return top
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	size := 0
	top := s.top
	for top != nil {
		size += 1
		top = top.GetNext()
	}
	return size
}
