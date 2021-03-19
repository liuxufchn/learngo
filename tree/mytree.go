package tree

type MyNode struct {
	*Node
}

func (node *MyNode) PostTraverse() {
	if node != nil && node.Node != nil {
		left := MyNode{node.Left}
		right := MyNode{node.Right}

		left.PostTraverse()
		right.PostTraverse()

		node.Print()
	}
}
