package main

import (
	"fmt"
)

// Create an expression tree using golang structs. For example string a + b - c on tree would look like the following,
//		  +
//		 / \
//	    a   -
//	  	   / \
//	 	  b   c
//
//
// Building the tree can be done manually through code. Once the tree is built, traverse the tree in preorder and postorder format and print the values.
//
//
// Expectation: Use Recursion, Strings, Pointers, Struct

type Node struct {
	val   byte
	left  *Node
	right *Node
}

func isOperator(ch byte) bool {
	if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
		return true
	}
	return false
}

func buildTree(str string, s int, e int) *Node {
	if s > e {
		return nil
	}

	opIdx := -1
	loPre := -1

	for i := s; i <= e; i++ {
		if isOperator(str[i]) {
			if str[i] == '+' || str[i] == '-' {
				opIdx = i
				loPre = 1
			} else {
				if loPre == -1 || loPre == 1 {
					opIdx = i
					loPre = 2
				}
			}
		}
	}
	if opIdx == -1 {
		return &Node{val: str[s]}
	}

	node := Node{val: str[opIdx]}
	node.left = buildTree(str, s, opIdx-1)
	node.right = buildTree(str, opIdx+1, e)

	return &node
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(string(root.val))
	printTree(root.left)
	printTree(root.right)
	//printTreeHelper(root, 0, "root")
}

//func printTreeHelper(root *Node, level int, side string) {
//	if root == nil {
//		return
//	}
//	// Print the current node with proper indentation
//
//	fmt.Printf("%s%s\n", strings.Repeat("  ", lvls-level), string(root.val))
//	// Recursively print left and right subtrees with increased level
//	if root.left != nil || root.right != nil {
//		printTreeHelper(root.left, level+1, "left")
//		printTreeHelper(root.right, level+1, "right")
//	}
//}

func printnPreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(string(root.val) + " ")
	printTree(root.left)
	printTree(root.right)
}

func printnPostOrder(root *Node) {
	if root == nil {
		return
	}
	printTree(root.left)
	printTree(root.right)
	fmt.Print(string(root.val) + " ")
}

func main() {
	str := "1+2-3"

	root := buildTree(str, 0, len(str)-1)
	//printTree(root)
	printnPreOrder(root)
	fmt.Println()
	printnPostOrder(root)
}
