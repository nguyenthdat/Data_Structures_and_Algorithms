//Doubly Linked Lists
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

func (l *DoublyLinkedList) addNodeAfter(data string, key string) {
	current_node := l.head
	var new_node Node
	for current_node != nil {
		if current_node.next == nil && current_node.data == key {
			l.append(data)
			return
		} else if current_node.data == key {
			new_node.data = data
			next_node := current_node.next
			current_node.next = &new_node
			new_node.next = next_node
			next_node.prev = &new_node
			new_node.prev = current_node
			return
		}
		current_node = current_node.next
	}
}

func (l *DoublyLinkedList) addNodeBefore(data string, key string) {
	current_node := l.head
	var new_node Node
	for current_node != nil {
		if current_node.prev == nil && current_node.data == key {
			l.prepend(data)
			return
		} else if current_node.data == key {
			new_node.data = data
			prev_node := current_node.prev
			prev_node.next = &new_node
			current_node.prev = &new_node
			new_node.next = current_node
			new_node.prev = prev_node
			return
		}
		current_node = current_node.next
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
	dlst.addNodeBefore("f", "b")
	dlst.addNodeBefore("h", "e")
	dlst.addNodeAfter("k", "d")
	dlst.print_list()
}
