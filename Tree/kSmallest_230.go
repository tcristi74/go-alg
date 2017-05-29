package BinaryTree

import "fmt"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

//Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
func kthSmallest(root *TreeNode, k int) int {
	idx:=0
	val :=-9999999
	traverseInOrder(root,k,idx,&val)
	return idx

}

func traverseInOrder(node *TreeNode, k int, idx int,val *int) {

	if node == nil {
		return
	}
	fmt.Printf("traversing node with value=%v,", node.Val)
	fmt.Println()

	// go left
	traverseInOrder(node.Left,k,idx,val)
	//*arr = append(*arr, node.Val)
	idx++
	if k==idx {
		*val= node.Val
//		return node.Val,true
	}
	traverseInOrder(node.Right,k,idx,val)
}

