package main

import (
	"fmt"
	"strconv"
	"strings"
)

// The type Stack  is a slice of integers ([]int). It allows you to create a stack data structure that can hold a sequence of integers.
// Push(val int): Adds a value to the top of the stack.
// Pop() int: Removes and returns the top value from the stack.
// IsEmpty() bool: Checks if the stack is empty.
// Size() int: Returns the number of elements in the stack.
// Peek() int: Returns the top value of the stack without removing it.
// ReverseStack(): Reverses the order of elements in the stack.
// Values() []int: Returns a slice of all elements in the stack.
// String() string: Returns a string representation of the stack.
// IterativeStack(): Prints the values of all elements in the stack.
// Copy() *Stack: Creates a copy of the stack.
// IsEqual(other *Stack) bool: Checks if two stacks have the same elements in the same order.
type Stack []int

// NewStack creates a new instance of the Stack type and returns a pointer to it.
//
// No parameters.
// Returns:
// - *Stack: a pointer to a new instance of the Stack type.
func NewStack() *Stack {
	return &Stack{}
}

// Push appends a value to the stack.
//
// Parameters:
// - val: the integer value to be added to the stack.
// Return type: none.
func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// Pop removes and returns the top element from the stack.
//
// Returns:
// - int: the top element of the stack.
func (s *Stack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// IsEmpty checks if the stack is empty.
//
// It returns true if the stack is empty, false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the number of elements in the stack.
//
// It does not take any parameters.
// Return type: int.
func (s *Stack) Size() int {
	return len(*s)
}

// Peek returns the top element of the stack.
//
// It does not take any parameters.
// Return type: int.
func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

// ReverseStack reverses the stack by swapping elements.
//
// No parameters.
// No return type.
func (s *Stack) ReverseStack() {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i], (*s)[len(*s)-i-1] = (*s)[len(*s)-i-1], (*s)[i]
	}
}

// Values returns the elements of the stack as a slice of integers.
//
// No parameters.
// Return type: []int.
func (s *Stack) Values() []int {
	return *s
}

// String returns a string representation of the Stack.
//
// It iterates over the stack and converts each value to a string using the strconv.Itoa() function.
// The resulting strings are then joined together with a comma separator and enclosed in curly braces.
//
// Returns:
// - A string representation of the Stack.
func (s *Stack) String() string {
	values := make([]string, 0, len(*s))
	for _, val := range *s {
		values = append(values, strconv.Itoa(val))
	}
	return "{" + strings.Join(values, ", ") + "}"
}

// IterativeStack iterates over the elements of the stack and prints them.
//
// No parameters.
// No return type.
func (s *Stack) IterativeStack() {
	for _, val := range *s {
		fmt.Println(val)
	}
}

// Copy creates a copy of the Stack by pushing each element from the original stack onto a new stack.
//
// Returns:
// - A pointer to a new Stack containing the same elements as the original Stack.
func (s *Stack) Copy() *Stack {
	newStack := NewStack()
	for _, val := range *s {
		newStack.Push(val)
	}
	return newStack
}

// IsEqual checks if two stacks are equal by comparing their sizes and values.
//
// Parameters:
// - other: a pointer to the other Stack to compare equality with.
//
// Returns:
// - bool: true if the stacks are equal, false otherwise.
func (s *Stack) IsEqual(other *Stack) bool {
	if s.Size() != other.Size() {
		return false
	}
	for i := 0; i < s.Size(); i++ {
		if s.Values()[i] != other.Values()[i] {
			return false
		}
	}
	return true
}
