package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func main() {

	tree := NewBinarySearchTree()

	tree.Insert(50)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(51)
	tree.Insert(77)
	tree.Insert(65)
	tree.Insert(33)
	tree.Insert(99)

	// tree.InOrderTraverse()
	// tree.PreOrderTraverse()
	// tree.PostOrderTraverse()

	// println(*tree.MinNode())
	tree.BFS()
}

func (node *Node) Print() {
	fmt.Printf("Node address: %p, node Value: %v . Node Right: %p. Node Left :%p .\n", node, node.Val, node.Right, node.Left)
}

func (bst *BinarySearchTree) Insert(value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	var node *Node = &Node{Val: value}
	if bst.root == nil {
		bst.root = node
		return
	}
	bst.root.InsertNode(node)
}

func (node *Node) InsertNode(newNode *Node) {
	if node == nil {
		return
	}
	if node.Val > newNode.Val {
		if node.Left == nil {
			node.Left = newNode
			return
		}
		node.Left.InsertNode(newNode)
		return
	}
	if node.Val <= newNode.Val {
		if node.Right == nil {
			node.Right = newNode
			return
		}
		node.Right.InsertNode(newNode)
	}
}

func (bst *BinarySearchTree) InOrderTraverse() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.root.InOrderTraverseNode()
}

func (node *Node) InOrderTraverseNode() {
	if node != nil {
		node.Left.InOrderTraverseNode()
		node.Print()
		node.Right.InOrderTraverseNode()
	}
}

func (bst *BinarySearchTree) PreOrderTraverse() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.root.PreOrderTraverseNode()
}

func (node *Node) PreOrderTraverseNode() {
	if node != nil {
		node.Print()
		node.Left.PreOrderTraverseNode()
		node.Right.PreOrderTraverseNode()
	}
}

func (bst *BinarySearchTree) PostOrderTraverse() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.root.PostOrderTraverseNode()
}

func (node *Node) PostOrderTraverseNode() {
	if node != nil {
		node.Left.PostOrderTraverseNode()
		node.Right.PostOrderTraverseNode()
		node.Print()
	}
}

func (bst *BinarySearchTree) MinNode() *int {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	if bst.root == nil {
		return (*int)(nil)
	}
	return bst.root.MinNode()
}

func (node *Node) MinNode() *int {
	if node.Left == nil {
		return &node.Val
	}
	return node.Left.MinNode()
}

func (bst *BinarySearchTree) MaxNode() *int {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	if bst.root == nil {
		return (*int)(nil)
	}
	return bst.root.MaxNode()
}

func (node *Node) MaxNode() *int {
	if node.Right == nil {
		return &node.Val
	}
	return node.Right.MaxNode()
}

func (node *Node) IsLeaf() bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func (bst *BinarySearchTree) IsEmpty() bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return bst.root == nil
}

func (bst *BinarySearchTree) SearchValue(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return bst.root.SearchValueNode(key)
}

func (node *Node) SearchValueNode(key int) bool {
	if node == nil {
		return false
	}
	if node.Val > key {
		return node.Left.SearchValueNode(key)
	}
	if node.Val < key {
		return node.Right.SearchValueNode(key)
	}
	return true
}

func (bst *BinarySearchTree) Remove(value int) *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return bst.root.RemoveNode(value)
}

func (node *Node) RemoveNode(value int) *Node {

	if node == nil {
		return nil
	}

	if node.Val > value {
		node.Left = node.Left.RemoveNode(value)
		return node
	}

	if node.Val < value {
		node.Right = node.Right.RemoveNode(value)
		return node
	}

	if node.Left == nil && node.Right == nil {
		node = nil
		return nil
	}

	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}
	// minValue := node.Right.MinNode()
	// node.Val = *minValue
	// node.Right = node.Right.RemoveNode(*minValue)

	//inOrder predecessor
	var leftMostRightNode *Node
	leftMostRightNode = node.Right

	for {
		if leftMostRightNode == nil && leftMostRightNode.Left == nil {
			leftMostRightNode = leftMostRightNode.Left
			continue
		}
		break
	}
	node.Val = leftMostRightNode.Val
	node.Right = node.Right.RemoveNode(node.Val)

	//inOrder successor
	var rightMostRightNode *Node
	rightMostRightNode = node.Left
	for {
		if rightMostRightNode == nil && rightMostRightNode.Right == nil {
			rightMostRightNode = rightMostRightNode.Right
			continue
		}
		break
	}
	node.Val = rightMostRightNode.Val
	node.Right = node.Right.RemoveNode(node.Val)
	return node
}

func (bst *BinarySearchTree) String() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	bst.root.stringifyNode(0)
}

func (node *Node) stringifyNode(level int) {
	if node == nil {
		return
	}
	level++
	node.Right.stringifyNode(level)
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Println(node.Val)
	node.Left.stringifyNode(level)
}

func (bst *BinarySearchTree) DFS() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	nodes := bst.root.DFS()

	for _, node := range nodes {
		node.Print()
	}

}

func (node *Node) DFS() []*Node {
	visited := []*Node{}
	stack := []*Node{node}

	if node == nil {
		return visited
	}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		visited = append(visited, node)
		stack = stack[:len(stack)-1] //lifo

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	return visited
}

func (bst *BinarySearchTree) BFS() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	nodes := bst.root.BFS()
	for _, node := range nodes {
		node.Print()
	}
}

func (node *Node) BFS() []*Node {
	visited := []*Node{}
	queue := []*Node{node}

	if node == nil {
		return visited
	}

	for len(queue) != 0 {
		node := queue[0]
		visited = append(visited, node)
		queue = queue[1:] //fifo

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return visited
}
