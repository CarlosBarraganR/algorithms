package main

import (
	"fmt"
	"time"
)

type TreeNodeValue interface{} // specific interface definition IRL

type TreeNode struct {
	key   string
	value TreeNodeValue
	left  *TreeNode
	right *TreeNode
	count int
}

type BinaryTree struct {
	root *TreeNode
}

func (bt *BinaryTree) Search(key string) TreeNodeValue {
	// always start on root of the tree
	currentNode := bt.root
	// iterate until value is found or no more TreeNodes are available
	for currentNode != nil {
		if key < currentNode.key { // less moves to the left
			currentNode = currentNode.left
		} else if key > currentNode.key { // greater moves to the right
			currentNode = currentNode.right
		} else {
			return currentNode.value // if equal value found
		}
	}
	return nil
}

func (bt *BinaryTree) _count(node *TreeNode) int {
	// fmt.Println("node", node)
	if node == nil {
		return 0
	}
	return node.count
}

func (bt *BinaryTree) Size() int {
	return bt._count(bt.root)
}

func (bt *BinaryTree) _put(node *TreeNode, key string, value TreeNodeValue) *TreeNode {
	// recursive base condition to stop at the bottom of the tree
	if node == nil {
		return &TreeNode{key: key, value: value, count: 1}
	} else if key < node.key {
		node.left = bt._put(node.left, key, value)
	} else if key > node.key {
		node.right = bt._put(node.right, key, value)
	} else {
		node.value = value // overrides old value
	}

	node.count = 1 + bt._count(node.left) + bt._count(node.right)

	return node
}

func (bt *BinaryTree) Insert(key string, value TreeNodeValue) {
	// recursive way to insert and reset links all the way up
	bt.root = bt._put(bt.root, key, value)
}

func (bt *BinaryTree) _floor(node *TreeNode, key string) *TreeNode {
	// case 0: recursive base condition to stop when the bottom of the tree is reached
	if node == nil {
		return nil
	} else if key == node.key { // case 1: key equals the node key at root
		return node // the floor is current node
	} else if key < node.key { // case 2: key is less the node key at root
		return bt._floor(node.left, key) // the floor of key is in the left subtree
	}
	t := bt._floor(node.right, key) // case 3: key is greater than the node key at root
	if t != nil {
		return t // is there any key less or equal in the right subtree
	} else {
		return node // otherwise it is the root
	}
}

func (bt *BinaryTree) Floor(key string) *TreeNode {
	// recursive way to search and find key or closest (round down) value
	return bt._floor(bt.root, key)
}

func (bt *BinaryTree) _rank(node *TreeNode, key string) int {
	// recursive base condition to stop when no keys are less than key
	if node == nil {
		return 0
	} else if key < node.key { // if key is less we rank the next node to the left
		return bt._rank(node.left, key)
	} else if key > node.key { // if greater means we need to count current node childrens
		return 1 + bt._count(node.left) + bt._count(node.right)
	} else { // if equal we count to the left only since right will always be greater
		return bt._count(node.left)
	}
}

func (bt *BinaryTree) Rank(key string) int {
	// recursive way to find how many keys are less than the key
	return bt._rank(bt.root, key)
}

func (bt *BinaryTree) _deleteMin(node *TreeNode) *TreeNode {
	// if min is found return right node that will replace deleted min
	if node.left == nil {
		return node.right
	}
	node.left = bt._deleteMin(node.left)                          // if min is found will be replaced by the right node
	node.count = 1 + bt._count(node.left) + bt._count(node.right) // update node count to adjust deleted min node
	return node

}

func (bt *BinaryTree) DeleteMin() {
	// recursive way to find how many keys are less than the key
	bt.root = bt._deleteMin(bt.root)
}

func main() {
	start := time.Now()

	binaryTree := &BinaryTree{}

	binaryTree.Insert("s", "S")
	binaryTree.Insert("e", "E")
	binaryTree.Insert("a", "A")
	binaryTree.Insert("r", "R")
	binaryTree.Insert("c", "C")
	binaryTree.Insert("h", "H")
	binaryTree.Insert("m", "M")
	binaryTree.Insert("x", "X")

	fmt.Println(binaryTree.Search("h")) // H
	fmt.Println(binaryTree.Search("g")) // null
	fmt.Println(binaryTree.Floor("g"))  // E
	fmt.Println(binaryTree.Size())      // 8
	binaryTree.Insert("g", "G")
	fmt.Println(binaryTree.Search("g")) // G
	binaryTree.Insert("g", "new G")
	fmt.Println(binaryTree.Search("g")) // new G
	fmt.Println(binaryTree.Size())      // 9
	fmt.Println(binaryTree.Rank("e"))   // 2
	binaryTree.DeleteMin()
	fmt.Println(binaryTree.Size())      // 8
	fmt.Println(binaryTree.Search("a")) // nil because it was deleted as the minimum value

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
