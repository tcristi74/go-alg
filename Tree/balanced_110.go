package BinaryTree

import "fmt"

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	max := 0
	min := -1
	traverse(root, 0, &min, &max)
	ret := true
	if min != -1 && min < max-1 {
		fmt.Println("result is false")
		ret = false
	}

	fmt.Printf("min=%d\n", min)
	fmt.Printf("max=%d\n", max)
	return ret
}

func traverse(node *TreeNode, current int, min *int, max *int) {

	fmt.Printf("null node value =%d current=%d\n", node.Val, current)
	if node.Left == nil || node.Right == nil {
		fmt.Printf("null node value =%d\n", node.Val)
		if *min > current || *min == -1 {
			*min = current
			fmt.Printf("get min=%d\n", *min)
		}
	}

	//fmt.Println("aaa")
	if node.Left != nil {
		if *max < current+1 {
			*max = current + 1
		}
		traverse(node.Left, current+1, min, max)
	}
	if node.Right != nil {
		if *max < current+1 {
			*max = current + 1
		}
		traverse(node.Right, current+1, min, max)
	}
}
