package main

import "fmt"

type node struct {
	left  *node
	right *node
	data  int64
}

type binaryTree struct {
	root *node
}

func main() {

	root := &node{}
	fmt.Println(root.data)

}

func (root *node) addNode(newNode *node) {

	switch {
	case newNode == nil:
		panic("the newNode parameter of addNode function can't be nil")
	case newNode.data > root.data:
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.addNode(newNode)
		}
	case newNode.data <= root.data:
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.addNode(newNode)
		}
	}

}

// func (bt *binaryTree) addNode(data int64) {
//
// 	newNode := &node{data: data}
//
// 	baseNode := &bt.root
//
// 	for {
// 		switch {
// 		case (*baseNode) == nil:
// 			*baseNode = newNode
// 			return
// 		case (*baseNode).data <= newNode.data:
// 			baseNode = &(*baseNode).left
// 		case (*baseNode).data > newNode.data:
// 			//			**baseNode = *(*baseNode).right
// 			return
// 		default:
// 			return
// 		}
// 	}
// }
