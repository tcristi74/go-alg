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


}





