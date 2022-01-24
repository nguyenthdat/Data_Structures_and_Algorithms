//Doubly Linked Lists
package main

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	data int
}

type DoublyLinkedList struct {
	head *Node
	//tail *Node
}

func (l *DoublyLinkedList) append(data int) {
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

func (l *DoublyLinkedList) prepend(data int) {
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

func (l *DoublyLinkedList) addNodeAfter(data int, key int) {
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

func (l *DoublyLinkedList) addNodeBefore(data int, key int) {
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

func (l *DoublyLinkedList) delete(key int) {
	current_node := l.head
	for current_node != nil {
		if current_node.data == key && current_node == l.head {
			// Case 1
			if current_node.next == nil {
				current_node = nil
				l.head = nil
				return
			} else { // Case 2
				next_node := current_node.next
				current_node.next = nil
				current_node = nil
				next_node.prev = nil
				l.head = next_node
				return
			}
		} else if current_node.data == key {
			if current_node.next != nil { // Case 3
				next_node := current_node.next
				prev_node := current_node.prev
				prev_node.next = next_node
				next_node.prev = prev_node
				current_node.prev, current_node.next, current_node = nil, nil, nil
				return
			} else { // Case 4
				prev_node := current_node.prev
				prev_node.next = nil
				current_node.prev = nil
				current_node = nil
				return
			}
		}
		current_node = current_node.next
	}
}
func (l *DoublyLinkedList) deleteNode(node *Node) {
	current_node := l.head
	for current_node != nil {
		if current_node == node && current_node == l.head {
			// Case 1
			if current_node.next == nil {
				current_node = nil
				l.head = nil
				return
			} else { // Case 2
				next_node := current_node.next
				current_node.next = nil
				current_node = nil
				next_node.prev = nil
				l.head = next_node
				return
			}
		} else if current_node == node {
			if current_node.next != nil { // Case 3
				next_node := current_node.next
				prev_node := current_node.prev
				prev_node.next = next_node
				next_node.prev = prev_node
				current_node.prev, current_node.next, current_node = nil, nil, nil
				return
			} else { // Case 4
				prev_node := current_node.prev
				prev_node.next = nil
				current_node.prev = nil
				current_node = nil
				return
			}
		}
		current_node = current_node.next
	}
}

func (l *DoublyLinkedList) reverse() {
	current_node := l.head
	var tmp_node *Node
	for current_node != nil {
		tmp_node = current_node.prev
		current_node.prev = current_node.next
		current_node.next = tmp_node
		current_node = current_node.prev
	}
	if tmp_node != nil {
		l.head = tmp_node.prev
	}
}

func (l *DoublyLinkedList) removeDuplicates() {
	current_node := l.head
	seen := make(map[int]int)
	for current_node != nil {
		if _, ok := seen[current_node.data]; !ok {
			seen[current_node.data] = 1
			current_node = current_node.next
		} else {
			next_node := current_node.next
			l.deleteNode(current_node)
			current_node = next_node
		}
	}
}

func (l *DoublyLinkedList) pairsWithSum(sum int) []string {
	var pairs []string
	right := l.head
	var left *Node
	for right != nil {
		left = right.next
		for left != nil {
			if right.data+left.data == sum {
				pairs = append(pairs, fmt.Sprintf("(%d, %d)", right.data, left.data))
			}
			left = left.next
		}
		right = right.next
	}
	return pairs
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
	dlst.append(1)
	dlst.append(2)
	dlst.append(3)
	dlst.append(4)
	dlst.append(5)
	fmt.Println(dlst.pairsWithSum(5))

}
