package main

import (
	"fmt"
)

type stack struct {
	// Stack structure
	stackData []interface{};
}

func (s *stack) push(data ...interface{}) {
	// Pushes data into the stack
	for _, i := range data {
		s.stackData = append(s.stackData, i);
	}
}

func (s *stack) pop() interface{} {
	// Pops and returns element at the top of the stack
	var top interface{} = s.stackData[len(s.stackData) - 1];
	s.stackData = s.stackData[:len(s.stackData) - 1];

	return top;		// Return top of stack
}

func (s *stack) top() interface{} {
	// Returns element at the top of the stack
	return s.stackData[len(s.stackData) - 1];
}

func (s stack) display() {
	// Display stack elements from top to bottom
	for i := len(s.stackData) - 1; i >= 0; i-- {
		fmt.Printf("%v ", s.stackData[i]); // Display data
		fmt.Println();
	}
}
