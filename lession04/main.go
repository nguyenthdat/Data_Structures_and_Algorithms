//Binary Trees in GoLang

/*
Binary Tree
A binary tree is a tree data structure in which each node
has at most two children, which are referred to as the left
child and the right child.
*/

package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

type BinaryTree struct {
	Root *Node
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d - ", root.Data)
	preOrder(root.Left)
	preOrder(root.Right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Printf("%d - ", root.Data)
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%d - ", root.Data)
	inOrder(root.Right)
}

func main() {
	var tree BinaryTree
	tree.Root = &Node{nil, nil, 1}
	tree.Root.Left = &Node{nil, nil, 2}
	tree.Root.Right = &Node{nil, nil, 3}
	tree.Root.Left.Left = &Node{nil, nil, 4}
	tree.Root.Left.Right = &Node{nil, nil, 5}
	tree.Root.Right.Left = &Node{nil, nil, 6}
	tree.Root.Right.Right = &Node{nil, nil, 7}
	preOrder(tree.Root)
	fmt.Println()
	postOrder(tree.Root)
	fmt.Println()
	inOrder(tree.Root)
}
