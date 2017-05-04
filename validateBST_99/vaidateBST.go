package validBst

import "fmt"

type TreeNode struct {
	Val   int`2`
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
