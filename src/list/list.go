package list

import (
	"github.com/avddvd/algos-ds/node"
	"reflect"
)

type List struct {
	head *node.Node
}

func (l *List) Push(data interface{}) {
	node := new(node.Node)
	node.SetData(data)
	if l.head == nil {
		l.head = node
	} else {
		node.SetNext(l.head)
		l.head = node
	}
}

func (l *List) Append(data interface{}) {
	node := new(node.Node)
	node.SetData(data)
	if l.head == nil {
		l.head = node
	} else {
		curr := l.head
		for curr.GetNext() != nil {
			curr = curr.GetNext()
		}
		curr.SetNext(node)
	}
}

func (l *List) Pop() *node.Node {
	curr := l.head
	if curr == nil {
		return nil
	}
	var prev *node.Node
	for curr != nil && curr.GetNext() != nil {
		prev = curr
		curr = curr.GetNext()
	}
	if prev == nil {
		l.head = nil
	} else {
		prev.SetNext(nil)
	}
	return curr
}

func (l *List) PopHead() *node.Node {
	var top *node.Node
	if l.head != nil {
		top = l.head
		l.head = top.GetNext()
	}
	return top
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Size() int {
	sum := 0
	curr := l.head
	for curr != nil {
		sum += 1
		curr = curr.GetNext()
	}
	return sum
}

func (l *List) Search(data interface{}) bool {
	curr := l.head
	exists := false
	for curr != nil && !exists {
		if reflect.DeepEqual(curr.GetData(), data) {
			exists = true
		}
		curr = curr.GetNext()
	}
	return exists
}

func (l *List) Index(data interface{}) int {
	curr := l.head
	index := -1
	exists := false
	for curr != nil && !exists {
		if reflect.DeepEqual(curr.GetData(), data) {
			exists = true
		}
		index += 1
		curr = curr.GetNext()
	}
	return index
}

func (l *List) Remove(data interface{}) {
	var prev *node.Node
	curr := l.head
	found := false
	for curr != nil && !found {
		if reflect.DeepEqual(curr.GetData(), data) {
			found = true
		} else {
			prev = curr
			curr = curr.GetNext()
		}
	}
	if prev == nil {
		if curr == nil {
			return
		} else {
			l.head = curr.GetNext()
		}
	} else {
		prev.SetNext(curr.GetNext())
	}
}

func (l *List) Insert(pos int, data interface{}) {
	var prev *node.Node
	curr := l.head
	index := 0
	node := new(node.Node)
	node.SetData(data)
	for index < pos && curr != nil {
		prev = curr
		curr = curr.GetNext()
		index += 1
	}
	if prev == nil {
		l.head = node
	} else {
		node.SetNext(curr)
		prev.SetNext(node)
	}
}

func (l *List) Delete(pos int) *node.Node {
	index := 0
	if pos < 0 || pos >= l.Size() {
		return nil
	}
	var prev *node.Node
	var res *node.Node
	curr := l.head
	for index < pos && curr != nil {
		index += 1
		prev = curr
		curr = curr.GetNext()
	}
	if curr == nil {
		return nil
	}
	if prev == nil {
		res = curr
		l.head = nil
	} else {
		prev.SetNext(curr.GetNext())
		curr.SetNext(nil)
		res = curr
	}
	return res
}

func (l *List) Reverse() {
	curr := l.head
	if curr == nil {
		return
	}
	var prev *node.Node
	var next *node.Node
	for curr != nil {
		next = curr.GetNext()
		curr.SetNext(prev)
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *List) String() string {
	curr := l.head
	r := ""
	for curr != nil {
		r = r + "->" + string(curr.GetData().(rune))
		curr = curr.GetNext()
	}
	return r
}

func (l *List) Split() (*List, *List) {
	slow := l.head
	var fast *node.Node
	l1, l2 := new(List), new(List)
	if slow != nil {
		fast = l.head.GetNext()
	}
	for fast != nil {
		slow = slow.GetNext()
		if fast.GetNext() != nil {
			fast = fast.GetNext().GetNext()
		}
	}
	l2.head = slow.GetNext()
	slow.SetNext(nil)
	l1.head = l.head
	l.head = nil
	return l1, l2
}
