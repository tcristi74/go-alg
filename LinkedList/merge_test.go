package LinkedList

import (
	"fmt"
	"testing"
)


func TestSortedList(t *testing.T) {

	n1 := ListNode{1, nil}
	n1.Next = &ListNode{3, nil}
	n1.Next.Next = &ListNode{7, nil}
	n1.Next.Next.Next = &ListNode{11, nil}
	n1.Next.Next.Next.Next = &ListNode{13, nil}

	n2 := ListNode{2, nil}
	n2.Next = &ListNode{4, nil}
	n2.Next.Next = &ListNode{10, nil}


	res := MergeTwoLists(&n1, &n2)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestSorted2List(t *testing.T) {

	n1 := ListNode{1, nil}
	//n1.Next = &ListNode{3, nil}
	//n1.Next.Next = &ListNode{7, nil}
	//n1.Next.Next.Next = &ListNode{11, nil}
	//n1.Next.Next.Next.Next = &ListNode{13, nil}

	n2 := ListNode{2, nil}
	//n2.Next = &ListNode{4, nil}
	//n2.Next.Next = &ListNode{10, nil}


	res := MergeTwoLists(&n1, &n2)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}


func TestSortedList2(t *testing.T) {

	n1 := ListNode{1, nil}
	//n1.Next = &ListNode{3, nil}
	//n1.Next.Next = &ListNode{7, nil}
	//n1.Next.Next.Next = &ListNode{11, nil}
	//n1.Next.Next.Next.Next = &ListNode{13, nil}

	n2 := ListNode{2, nil}
	//n2.Next = &ListNode{4, nil}
	//n2.Next.Next = &ListNode{10, nil}

	fmt.Println(n1)
	fmt.Println(n2)


	res := MergeTwoLists(&n1, &n2)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestSortedList3(t *testing.T) {

	n1 := ListNode{1, nil}
	n1.Next = &ListNode{3, nil}
	n1.Next.Next = &ListNode{7, nil}
	//n1.Next.Next.Next = &ListNode{11, nil}
	//n1.Next.Next.Next.Next = &ListNode{13, nil}

	n2 := ListNode{0, nil}
	//n2.Next = &ListNode{4, nil}
	//n2.Next.Next = &ListNode{10, nil}

	fmt.Println(n1)
	fmt.Println(n2)


	res := MergeTwoLists(&n1, &n2)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}