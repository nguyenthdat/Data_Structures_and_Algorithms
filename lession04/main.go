//Binary Trees in GoLang

/*
Binary Tree
A binary tree is a tree data structure in which each node
has at most two children, which are referred to as the left
child and the right child.
*/

package main

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

type BinaryTree struct {
	Root *Node
}

var new_node Node

func initNode(data int) *Node {
	new_node.Data = data
	return &new_node
}

func main() {
	var tree BinaryTree
	tree.Root = initNode(1)
	tree.Root.Left = initNode(2)
	tree.Root.Right = initNode(3)
	tree.Root.Left.Left = initNode(4)
	tree.Root.Left.Right = initNode(5)
	tree.Root.Right.Left = initNode(6)
	tree.Root.Right.Right = initNode(7)
	tree.Root.Right.Right.Right = initNode(8)

}
