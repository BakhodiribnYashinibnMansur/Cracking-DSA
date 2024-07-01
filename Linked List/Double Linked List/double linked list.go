package main

// This struct definition represents a node in a linked list data structure. It has three fields: Val to store the value of the node, Prev to store the previous node in the list, and Next to store the next node in the list.
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

// LinkedList struct definition represents a linked list with head and tail pointers.
// AddToHead(val int): Adds a new node with the given value to the head of the linked list.
// AddToTailWithTail(val int): Adds a new node to the end of the linked list using the tail pointer.
// AddToTailWithoutTail(val int): Adds a new node to the end of the linked list without using the tail pointer.
// NodeBetWeen(left, right int): Finds the node between two given values in the linked list.
// IterativeList(): Iterates over the linked list and prints the values of each node.
// ReverseList(): Reverses the linked list.
type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) AddToHead(val int) {
	// AddToHead adds a new node with the given value to the head of the linked list.
	//
	// Parameters:
	// - val: the value to be added to the linked list.
	//
	// Return type: none.
	node := &Node{Val: val}
	if ll.head != nil {
		node.Next = ll.head
		ll.head.Prev = node
	}
	ll.head = node
}

// AddToTailWithTail adds a new node to the end of the linked list using the tail pointer.
//
// Parameters:
// - val: the value of the new node.
//
// Return type: none.
func (ll *LinkedList) AddToTailWithTail(val int) {
	node := &Node{Val: val}
	if ll.tail != nil {
		node.Prev = ll.tail
		ll.tail.Next = node
	}
	ll.tail = node
	// AddToTailWithoutTail adds a new node to the end of the linked list without using the tail pointer.
	//
	// Parameters:
	// - val: the value of the new node.
	//
	// Return type: none.
}
func (ll *LinkedList) AddToTailWithoutTail(val int) {
	node := &Node{Val: val}
	for currNode := ll.head; currNode.Next != nil; currNode = currNode.Next {
		if currNode.Next == nil {
			currNode.Next = node
			node.Prev = currNode
			break
		}
	}
}

// NodeBetween finds the node between two given values in the linked list.
//
// Parameters:
// - left: the value of the node to the left of the target node.
// - right: the value of the node to the right of the target node.
//
// Returns:
// - *Node: the node between the two given values, or nil if not found.
func (ll *LinkedList) NodeBetWeen(left, right int) *Node {
	var node *Node
	var nodeWith *Node
	for node = ll.head; node != nil; node = node.Next {
		if node.Prev != nil && node.Next != nil {
			if node.Prev.Val == left && node.Next.Val == right {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// IterativeList iterates over the linked list and prints the values of each node.
//
// No parameters.
// No return type.
func (ll *LinkedList) IterativeList() {
	for node := ll.head; node != nil; node = node.Next {
		println(node.Val)
	}
}

// ReverseList reverses the linked list.
//
// No parameters.
// No return type.

func (ll *LinkedList) ReverseList() {
	var prev *Node
	var curr *Node
	var next *Node
	curr = ll.head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
}
