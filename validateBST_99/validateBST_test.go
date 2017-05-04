package validBst

import (
	"fmt"
	"testing"
)


func ATestIsValidBST(t *testing.T) {

	root := TreeNode{Val: 20, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 18, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 36, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 19, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 26, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 46, Left: nil, Right: nil}

	root.Right.Left.Left = &TreeNode{Val: 24, Left: nil, Right: nil}
	root.Right.Left.Right = &TreeNode{Val: 27, Left: nil, Right: nil}

	root.Right.Right.Right = &TreeNode{Val: 100, Left: nil, Right: nil}

	fmt.Println("done loading the tree")

	if IsValidBST(&root) {
		fmt.Println("it is BST")
	} else {
		fmt.Println("it is NOT BST")
	}
}

//TestArrayToBst
func TestArrayToBst(t *testing.T) {

	fmt.Println("Start TestArrayToBst")
	arr:=[]int {3,6,1,3,7,2,89,4,45,12,10,34,5}
	fmt.Println(arr)
	tree:= ArrayToBST(arr)

	if IsValidBST(tree) {
		fmt.Println("it is BST")
	} else {
		fmt.Println("it is NOT BST")
	}
}
