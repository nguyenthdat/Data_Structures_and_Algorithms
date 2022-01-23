package main

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	data string
}

type DoublyLinkedList struct {
	head *Node
	//tail *Node
}

func (l *DoublyLinkedList) append(data string) {
	var new_node Node
	if l.head == nil {
		new_node.data = data
		new_node.prev = nil
		l.head = &new_node
		//l.tail = &new_node
	} else {
		new_node.data = data
		current_node := l.head
		for current_node.next != nil {
			current_node = current_node.next
		}
		current_node.next = &new_node
		new_node.prev = current_node
		new_node.next = nil
	}
}

func (l *DoublyLinkedList) prepend(data string) {
	var new_node Node
	if l.head == nil {
		new_node.data = data
		new_node.prev = nil
		l.head = &new_node
	} else {
		new_node.data = data
		l.head.prev = &new_node
		new_node.next = l.head
		l.head = &new_node
		new_node.prev = nil
	}
}

func (l *DoublyLinkedList) print_list() {
	current_node := l.head
	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next
	}

}

func main() {
	var dlst DoublyLinkedList
	dlst.append("a")
	dlst.append("b")
	dlst.append("c")
	dlst.append("d")
	dlst.prepend("e")
	dlst.print_list()
}
