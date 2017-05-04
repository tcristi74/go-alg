package main

import (
	"fmt"
)

import (
	//"github.com/tcristi74/go-alg/tests"
	//"github.com/tcristi74/go-alg/algorithm"
	//"github.com/tcristi74/go-alg/palindrome"
 	"github.com/tcristi74/go-alg/validateBST_99"
)


func init() {

}


func main() {

	root := validBst.TreeNode{Val: 20, Left: nil, Right: nil}
	root.Left = &validBst.TreeNode{Val: 18, Left: nil, Right: nil}
	root.Right = &validBst.TreeNode{Val: 36, Left: nil, Right: nil}
	root.Left.Left = &validBst.TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &validBst.TreeNode{Val: 19, Left: nil, Right: nil}
	root.Right.Left = &validBst.TreeNode{Val: 26, Left: nil, Right: nil}
	root.Right.Right = &validBst.TreeNode{Val: 46, Left: nil, Right: nil}

	root.Right.Left.Left = &validBst.TreeNode{Val: 24, Left: nil, Right: nil}
	root.Right.Left.Right = &validBst.TreeNode{Val: 27, Left: nil, Right: nil}

	root.Right.Right.Right = &validBst.TreeNode{Val: 100, Left: nil, Right: nil}

	fmt.Println("done loading the tree")

	if validBst.IsValidBST(&root) {
		fmt.Println("it is BST")
	} else {
		fmt.Println("it is NOT BST")
	}

}





