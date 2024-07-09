package main

// Node struct defines a node structure for a linked list.
// Val represents the integer value stored in the node.
// Next is a pointer to the next node in the linked list.
type Node struct {
	Val  int
	Next *Node
}

// The LinkedList (main.go:10-12:2) class definition in Go defines a structure for a linked list. It has a single field head which is a pointer to a Node (main.go:111-124:2) struct. The LinkedList (main.go:10-12:2) struct does not have any methods defined in this snippet. However, it is likely that there are methods defined elsewhere in the codebase that operate on this struct.

// Here are some possible methods that could be defined for the LinkedList (main.go:10-12:2) struct:

// AddToHead(val int): Adds a new node with the given value at the beginning of the linked list.
// AddToTail(val int): Adds a new node with the given value at the end of the linked list.
// IterateList(): Iterates through the linked list and prints information about each node.
// LastNode(): Returns the last node in the linked list.
// SearchNode(val int): Searches for a node with the given value in the linked list.
// AddAfterNode(nodeVal, val int): Adds a new node with the given value after a specific node value in the linked list.
// AddBeforeNode(nodeVal, val int): Adds a new node with the given value before a specific node value in the linked list.
// DeleteNode(nodeVal int): Deletes a node with the given value from the linked list.
// Slice(num int): Returns the node in the linked list that has the next node with the given value.
// Reverse(): Reverses the order of the linked list.
// These are just some examples, and there may be other methods defined depending on the specific requirements of the linked list implementation.
type LinkedList struct {
	head *Node
}

// Init initializes a new LinkedList and returns a pointer to it.
//
// No parameters.
// Returns a pointer to a LinkedList.

func Init() *LinkedList {
	return &LinkedList{}
}

// main is the entry point of the Go program.
//
// It initializes a linked list, adds numbers to the head of the list,
// adds a number to the tail of the list, adds a number after a specific
// node in the list, adds a number before a specific node in the list,
// deletes a node from the list, iterates through the list and prints
// information about each node, searches for a specific node in the
// list, slices the list, iterates through the list again, and reverses
// the list.
//
// There are no parameters.
// There is no return value.

func (node *Node) PrintNode() {
	if node == nil {
		println("node is nil")
		return
	}
	println("node's address is", node, "and its value is", node.Val)
}

// AddToHead adds a new node with the given value at the beginning of the linked list.
//
// Parameters:
// - val: an integer representing the value to be added to the linked list.
// Return type: None.
func (ll *LinkedList) AddToHead(val int) {
	head := &Node{Val: val}
	if ll.head != nil {
		head.Next = ll.head
	}
	ll.head = head
}

// AddToTail adds a new node with the given value at the end of the linked list.
//
// Parameters:
// - val: an integer representing the value to be added to the linked list.
//
// Return type: None.
func (ll *LinkedList) AddToTail(val int) {
	tail := &Node{Val: val}
	node := ll.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = tail
}

// IterateList iterates through the linked list and prints information about each node.
//
// No parameters.
// No return value.
func (ll *LinkedList) IterateList() {
	order := 0
	for node := ll.head; node != nil; node = node.Next {
		order++
		println("linkedList's order is", order, " . node's address is", node, "and its value is ", node.Val)
	}

}

// LastNode returns the last node in the linked list.
//
// Parameters:
// - ll: a pointer to a LinkedList struct.
//
// Return type:
// - *Node: a pointer to the last Node in the linked list.
func (ll *LinkedList) LastNode() *Node {
	node := ll.head
	for node.Next != nil {
		node = node.Next
	}
	return node
}

// SearchNode searches for a node with the given value in the linked list.
//
// Parameters:
// - val: an integer representing the value to search for.
//
// Return type:
// - *Node: a pointer to the first node with the given value, or nil if not found.
func (ll *LinkedList) SearchNode(val int) *Node {
	node := ll.head
	for node != nil {
		if node.Val == val {
			return node
		}
		node = node.Next
	}
	return nil
}

// AddAfterNode adds a new node with the given value after a specific node value in the linked list.
//
// Parameters:
// - nodeVal: an integer representing the value of the node after which the new node should be added.
// - val: an integer representing the value to be added to the linked list.
// Return type: a boolean indicating whether the operation was successful.
func (ll *LinkedList) AddAfterNode(nodeVal, val int) bool {
	node := ll.SearchNode(nodeVal)
	if node == nil {
		return false
	}
	newNode := &Node{Val: val, Next: node.Next}
	node.Next = newNode
	return true
}

// AddBeforeNode adds a new node with the given value before a specific node value in the linked list.
//
// Parameters:
// - nodeVal: an integer representing the value of the node before which the new node should be added.
// - val: an integer representing the value to be added to the linked list.
// Return type: a boolean indicating whether the operation was successful.
func (ll *LinkedList) AddBeforeNode(nodeVal, val int) bool {
	node := ll.head
	isFind := false
	for node != nil {
		if node.Next.Val == nodeVal {
			isFind = true
			break
		}
		node = node.Next
	}
	if !isFind {
		return false
	}
	newNode := &Node{Val: val, Next: node.Next}
	node.Next = newNode
	return true
}

// DeleteNode deletes a node with the given value from the linked list.
//
// Parameters:
// - nodeVal: an integer representing the value of the node to be deleted.
//
// Return type:
// - bool: a boolean indicating whether the operation was successful.
func (ll *LinkedList) DeleteNode(nodeVal int) bool {
	node := ll.head
	for node != nil {
		if node.Next.Val == nodeVal {
			node.Next = node.Next.Next
			return true
		}
		node = node.Next
	}
	return false
}

// Reverse reverses the order of the linked list.
//
// This function takes no parameters.
// It does not return anything.
func (ll *LinkedList) Reverse() {

	var curr = ll.head
	var prev *Node = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
}

// Slice returns the node in the linked list that has the next node with the given value.
//
// Parameters:
// - val: an integer representing the value of the next node to be found.
//
// Return type:
// - *Node: a pointer to the node that has the next node with the given value, or nil if not found.
func (ll *LinkedList) Slice(num int) *Node {
	node := ll.head
	order := 0
	for node != nil {
		order++
		if order == num {
			return node
		}
		node = node.Next
	}
	return nil
}

func (ll *LinkedList) Merge(ll2 *LinkedList) *LinkedList {
	ll3 := Init()
	for node := ll.head; node != nil; node = node.Next {
		ll3.AddToTail(node.Val)
	}
	for node := ll2.head; node != nil; node = node.Next {
		ll3.AddToTail(node.Val)
	}
	return ll3
}
