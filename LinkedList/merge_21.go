package LinkedList

import (
//	"fmt"
//"math"
//"fmt"
)
import "math"

//import "math"

//import "fmt"

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := ListNode{Val: 0, Next: &ListNode{}}
	rt := head.Next

	for {
		if l1.Next == nil && l2.Next == nil {
			rt.Val = int(math.Min(float64(l1.Val), float64(l2.Val)))
			rt.Next = &ListNode{}
			rt = rt.Next
			rt.Val = int(math.Max(float64(l1.Val), float64(l2.Val)))
			break
		}
		if l1.Val > l2.Val {
			rt.Val = l2.Val
			if l2.Next == nil {
				rt.Next = l1
				break
			} else {
				l2 = l2.Next
			}

		} else {
			rt.Val = l1.Val
			if l1.Next == nil {
				rt.Next = l2
				break
			} else {
				l1 = l1.Next
			}
		}
		rt.Next = &ListNode{}
		rt = rt.Next
	}
	head = *head.Next
	return &head
}
