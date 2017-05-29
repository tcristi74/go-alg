package main

import (
	"fmt"
)

import (
	//"github.com/tcristi74/go-alg/tests"
	//"github.com/tcristi74/go-alg/algorithm"
	//"github.com/tcristi74/go-alg/palindrome"
 	"github.com/tcristi74/go-alg/Add2Numbers_2"
)



func init() {

}


func main() {

<<<<<<< HEAD
	fmt.Println("Start TestArrayToBst")
	arr:=[]int {3,6,1,3,7,2,89,4,45,12,10,34,5}
	fmt.Println(arr)
	tree:= validBst.ArrayToBST(arr)

	fmt.Println("tree is done")

	if validBst.IsValidBST(tree) {
		fmt.Println("it is BST")
	} else {
		fmt.Println("it is NOT BST")
	}


	//root := validBst.TreeNode{Val: 20, Left: nil, Right: nil}
	//root.Left = &validBst.TreeNode{Val: 18, Left: nil, Right: nil}
	//root.Right = &validBst.TreeNode{Val: 36, Left: nil, Right: nil}
	//root.Left.Left = &validBst.TreeNode{Val: 5, Left: nil, Right: nil}
	//root.Left.Right = &validBst.TreeNode{Val: 19, Left: nil, Right: nil}
	//root.Right.Left = &validBst.TreeNode{Val: 26, Left: nil, Right: nil}
	//root.Right.Right = &validBst.TreeNode{Val: 46, Left: nil, Right: nil}
	//
	//root.Right.Left.Left = &validBst.TreeNode{Val: 24, Left: nil, Right: nil}
	//root.Right.Left.Right = &validBst.TreeNode{Val: 27, Left: nil, Right: nil}
	//
	//root.Right.Right.Right = &validBst.TreeNode{Val: 100, Left: nil, Right: nil}
	//
	//fmt.Println("done loading the tree")
	//
	//if validBst.IsValidBST(&root) {
	//	fmt.Println("it is BST")
	//} else {
	//	fmt.Println("it is NOT BST")
	//}
=======




	//n1 :=  Add2Numbers_2.ListNode{2, nil}
	//
	//n2 := Add2Numbers_2.ListNode{8, nil}

	n1 :=  Add2Numbers_2.ListNode{1, nil}
	n1.Next = &Add2Numbers_2.ListNode{2, nil}
	//n1.Next.Next = &Add2Numbers_2.ListNode{3, nil}
	//n1.Next.Next.Next = &Add2Numbers_2.ListNode{4, nil}
	//n1.Next.Next.Next.Next = &Add2Numbers_2.ListNode{5, nil}


	//
	//n2 := Add2Numbers_2.ListNode{5, nil}
	//n2.Next = &Add2Numbers_2.ListNode{6, nil}
	//n2.Next.Next = &Add2Numbers_2.ListNode{4, nil}
	//n2.Next.Next.Next = &Add2Numbers_2.ListNode{2, nil}
	res := Add2Numbers_2.RemoveNthFromEndEntry(&n1, 1)
	for {
		if res==nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}

	//v:= InsertDelete.Constructor()
	//v.Insert(4)
	//v.Insert(6)
	//v.Insert(7)
	//v.Insert(9)
	//v.Insert(2)
	//v.Insert(4)
	//
	//fmt.Printf("remove 100 = %v",  v.Remove(100))
	//fmt.Println()
	//
	//fmt.Println("M=",v.M)
	//fmt.Println("P=",v.P)
	//
	//fmt.Printf("remove 7 = %v",  v.Remove(7))
	//fmt.Println()
	//
	//fmt.Println("M=",v.M)
	//fmt.Println("P=",v.P)
	//
	//fmt.Printf("remove 4 = %v",  v.Remove(4))
	//fmt.Println()
	//
	//fmt.Println("M=",v.M)
	//fmt.Println("P=",v.P)
	//
	//
	//fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()
	////fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()
	//fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()
	//fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()
	//fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()
	//fmt.Printf("result = %v",  v.GetRandom())
	//fmt.Println()

>>>>>>> 425d8152b3642e61f6ae703a6b52e6654e15f3cf

}





