package main

import (
	"fmt"
)

type node struct {
	// Linked list node structure
	data interface{};
	next *node;
}

func createLinkedList() node {
	// Create and return the head of a new linked list
	var head node = node{data: 0, next: nil};

	return head;		// Return head
}

func (n *node) insert(newData interface{}) {
	// Insert new node at end of the list
	var newNode node = node{data: newData, next: nil};

	if n.next == nil {
		// If list is currently empty, insert at head
		n.next = &newNode;
	} else {
		var current *node;
		
		for current = n; current.next != nil; current = current.next {
			// Traverse the list
		}
		// Insert the new node
		current.next = &newNode;
	}
}

func (n *node) display() {
	// Display elements of a linked list
	for current := n.next; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.data);
	}
	fmt.Printf("END");	// Signify end of the list
	fmt.Println();
}
