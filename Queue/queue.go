package main

import "fmt"

// This code defines a new type called Queue in Go. It is a slice of integers.
// Add(val int): Appends a value to the queue.
// Peek() int: Returns the first element of the queue without removing it.
// Clear(): Clears the queue by resetting its length to 0.
// Copy() *Queue: Creates a new Queue by copying all the elements from the current Queue.
// Values() []int: Returns the elements of the queue.
// String() string: Returns a string representation of the Queue.
// IterativeQueue(): Iterates over the elements of the queue and prints them.
// ReverseQueue(): Reverses the elements of the queue in-place.
// IsFull() bool: Checks if the queue is full.
// Dequeue() int: Removes and returns the first element from the Queue.
// IsEmpty() bool: Checks if the queue is empty.
type Queue []int

// NewQueue creates and returns a new Queue.
//
// No parameters.
// Returns a pointer to the new Queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Add appends a value to the queue.
//
// Parameters:
// - val: the integer value to be added to the queue.
// Return type: none.
func (q *Queue) Add(val int) {
	*q = append(*q, val)
}

// Peek returns the first element of the queue without removing it.
//
// Returns:
// - int: the first element of the queue.
func (q *Queue) Peek() int {
	return (*q)[0]
}

// Clear clears the queue by resetting its length to 0.
//
// This function takes a pointer to a Queue as its receiver.
// It modifies the underlying slice to have a length of 0, effectively clearing the queue.
// There are no parameters.
// There is no return type.
func (q *Queue) Clear() {
	*q = (*q)[0:0]
}

// Copy creates a new Queue by copying all the elements from the current Queue.
//
// No parameters.
// Returns a pointer to a new Queue.
func (q *Queue) Copy() *Queue {
	newQueue := NewQueue()
	for _, val := range *q {
		newQueue.Add(val)
	}
	return newQueue
}

// Values returns the elements of the queue.
//
// No parameters.
// Return type: []int.
func (q *Queue) Values() []int {
	return *q
}

// String returns a string representation of the Queue.
//
// No parameters.
// Return type: string.
func (q *Queue) String() string {
	return fmt.Sprintf("%v", *q)
}

// IterativeQueue iterates over the elements of the queue and prints them.
//
// No parameters.
// No return type.
func (q *Queue) IterativeQueue() {
	for _, val := range *q {
		fmt.Println(val)
	}
}

// ReverseQueue reverses the elements of the queue in-place.
//
// No parameters.
// No return type.
func (q *Queue) ReverseQueue() {
	for i := 0; i < len(*q)/2; i++ {
		(*q)[i], (*q)[len(*q)-i-1] = (*q)[len(*q)-i-1], (*q)[i]
	}
}

// IsFull checks if the queue is full.
//
// No parameters.
// Return type: bool.
func (q *Queue) IsFull() bool {
	return false
}

// Dequeue removes and returns the first element from the Queue.
//
// Returns:
// - int: the value of the removed element.
func (q *Queue) Dequeue() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

// IsEmpty checks if the queue is empty.
//
// No parameters.
// Returns a boolean indicating whether the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Size returns the number of elements in the queue.
//
// No parameters.
// Return type: int.
func (q *Queue) Size() int {
	return len(*q)
}
