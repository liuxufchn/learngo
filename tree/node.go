package tree

import "fmt"

type Node struct {
	Value       int
	Right, Left *Node
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) PreTraverse() {
	if node != nil {
		node.Print()
		node.Left.PreTraverse()
		node.Right.PreTraverse()
	}
}
