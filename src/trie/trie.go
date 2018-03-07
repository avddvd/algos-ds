package trie

import "github.com/avddvd/algos-ds/stack"

type Node struct {
	value    string
	children map[string]*Node
}

func NewNode() *Node {
	return &Node{
		value:    "",
		children: map[string]*Node{},
	}
}

func (n *Node) GetValue() string {
	return n.value
}

func (n *Node) GetChildren() map[string]*Node {
	return n.children
}

func (n *Node) SetValue(value string) {
	n.value = value
}

func (n *Node) Find(key string) string {
	node := n
	for _, ch := range key {
		s := string(ch)
		if child, ok := node.GetChildren()[s]; ok {
			node = child
		} else {
			return ""
		}
	}
	return node.GetValue()
}

func (n *Node) FindPrefixMatches(key string) []string {
	results := []string{}
	node := n
	for _, ch := range key {
		s := string(ch)
		if child, ok := node.GetChildren()[s]; ok {
			node = child
		} else {
			return results
		}
	}
	s := new(stack.Stack)
	s.Push(node)
	for !s.IsEmpty() {
		node := s.Pop().GetData().(*Node)
		if node.GetValue() != "" {
			results = append(results, node.GetValue())
		}
		for _, v := range node.GetChildren() {
			s.Push(v)
		}
	}
	return results
}

func (n *Node) Insert(key string) {
	node := n
	i := 0
	for i < len(key) {
		s := string(key[i])
		if child, ok := node.GetChildren()[s]; ok {
			i += 1
			node = child
		} else {
			break
		}
	}

	for i < len(key) {
		s := string(key[i])
		newNode := NewNode()
		node.GetChildren()[s] = newNode
		node = newNode
		i += 1
	}
	node.SetValue(key)
}
