package validBst

import (

	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func IsValidBST(root *TreeNode) bool {

	var arr []int
	// do an Order traverse see if is in order
	if root == nil {
		return false
	}

	fmt.Println("start traversing...")
	traverseInOrder(root, &arr)

	fmt.Println("end traversing...")
	fmt.Println(arr)

	prevVal := arr[0]

	fmt.Println(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i] > prevVal {
			prevVal = arr[i]
		} else {
			return false
		}
	}
	return true
}


func ArrayToBST(arr []int)  *TreeNode {

	rootNode :=TreeNode{ Left:nil,Right:nil}
	sort.Ints(arr)
	createBst(&rootNode,&arr)
	return &rootNode

}

func createBst(node *TreeNode, arr *[]int){

	//if node ==nil {
	//	node = &TreeNode{ Left: nil, Right: nil}
	//}
	if len(*arr)==1 {
		node.Val=(*arr)[0]
		return
	}
	//get the middle
	middle:=len(*arr)/2
	node.Val=(*arr)[middle-1]
	//split slice in two
	if middle>1 {
		sliceLeft:=(*arr)[0:middle-1]
		node.Left = &TreeNode{ Left: nil, Right: nil}
		createBst(node.Left, &sliceLeft)
	}
	sliceRight:=(*arr)[middle:]
	node.Right = &TreeNode{ Left: nil, Right: nil}
	createBst(node.Right,&sliceRight)
}



func traverseInOrder(node *TreeNode, arr *[]int) {


	if node == nil {
		return
	}
	fmt.Printf("traversing node with value=%v,", node.Val)
	fmt.Println()

	// go left
	traverseInOrder(node.Left, arr)
	*arr = append(*arr, node.Val)
	fmt.Println(arr)
	traverseInOrder(node.Right, arr)
}
