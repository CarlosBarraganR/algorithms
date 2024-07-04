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
}

type BinaryTree struct {
	root TreeNode
}

func (bt *BinaryTree) Search(key string) TreeNodeValue {
	// always start on root of the tree
	currentNode := &bt.root
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

func (bt *BinaryTree) _put(node *TreeNode, key string, value TreeNodeValue) *TreeNode {
	// recursive base condition to stop at the bottom of the tree
	if node == nil {
		return &TreeNode{key: key, value: value}
	}
	if key < node.key {
		node.left = bt._put(node.left, key, value)
	} else if key > node.key {
		node.right = bt._put(node.right, key, value)
	} else {
		node.value = value // overrides old value
	}
	return node
}

func (bt *BinaryTree) Insert(key string, value TreeNodeValue) {
	// recursive way to insert and reset links all the way up
	_ = bt._put(&bt.root, key, value)
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
	// recursive way to search and find key or closest value
	return bt._floor(&bt.root, key)
}

func main() {
	start := time.Now()

	binaryTree := &BinaryTree{TreeNode{key: "s", value: "S",
		left: &TreeNode{key: "e", value: "E",
			left: &TreeNode{key: "a", value: "A",
				right: &TreeNode{key: "c", value: "C"}},
			right: &TreeNode{key: "r", value: "R",
				left: &TreeNode{key: "h", value: "H",
					right: &TreeNode{key: "m", value: "M"}}}},
		right: &TreeNode{key: "x", value: "X"}}}

	fmt.Println(binaryTree.Search("h")) // H
	fmt.Println(binaryTree.Search("g")) // null
	fmt.Println(binaryTree.Floor("h"))  // E
	binaryTree.Insert("g", "G")
	fmt.Println(binaryTree.Search("g")) // G
	binaryTree.Insert("g", "new G")
	fmt.Println(binaryTree.Search("g")) // new G

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
