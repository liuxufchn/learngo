package main

import (
	"fmt"
	"go.uber.org/zap"
	"learngo/hello"
	"learngo/tree"
)

func main01() {
	max := hello.LongestSubStrWithoutRepeatingChars("hello")
	fmt.Println(max)
}

func main02() {
	root := tree.Node{}
	root.SetValue(5)
	root.Left = tree.CreateTreeNode(3)
	root.Right = &tree.Node{Value: 2}
	root.Left.Right = new(tree.Node)
	root.PreTraverse()
	fmt.Println()
	node := &tree.MyNode{Node: &root}
	node.PostTraverse()
}

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("zap test")
}
