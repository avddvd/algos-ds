package queue

import (
	"github.com/avddvd/algos-ds/node"
)

type Queue struct {
	head *node.Node
}

func (q *Queue) Enqueue(data interface{}) {
	node := new(node.Node)
	node.SetData(data)
	if q.head == nil {
		q.head = node
	} else {
		node.SetNext(q.head)
		q.head = node
	}
}

func (q *Queue) Dequeue() *node.Node {
	curr := q.head
	var prev *node.Node
	if curr == nil {
		return nil
	}
	for curr != nil && curr.GetNext() != nil {
		prev = curr
		curr = curr.GetNext()
	}
	if prev == nil {
		q.head = nil
	} else {
		prev.SetNext(nil)
	}
	return curr
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Size() int {
	sum := 0
	curr := q.head
	for curr != nil {
		sum += 1
		curr = curr.GetNext()
	}
	return sum
}
