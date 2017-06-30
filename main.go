package main

import (
	"fmt"
)

import (
	//"github.com/tcristi74/go-alg/tests"
	//"github.com/tcristi74/go-alg/algorithm"
	//"github.com/tcristi74/go-alg/palindrome"
 	//"github.com/tcristi74/go-alg/Add2Numbers_2"
	"github.com/tcristi74/go-alg/LinkedList"
)



func init() {

}


func main() {


	n1 :=  LinkedList.ListNode{1, nil}
	n1.Next = &LinkedList.ListNode{3, nil}
	n1.Next.Next = &LinkedList.ListNode{7, nil}
	//n1.Next.Next.Next = &ListNode{11, nil}
	//n1.Next.Next.Next.Next = &ListNode{13, nil}

	n2 := LinkedList.ListNode{0, nil}
	//n2.Next = &ListNode{4, nil}
	//n2.Next.Next = &ListNode{10, nil}

	fmt.Println(n1)
	fmt.Println(n2)


	res := LinkedList.MergeTwoLists(&n1, &n2)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}


}





