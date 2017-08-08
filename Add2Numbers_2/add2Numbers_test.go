package Add2Numbers_2

import (
	"fmt"
	"testing"
)

//TestAdd2NumbersEntry
func ATestAdd2NumbersEntry(t *testing.T) {
	n1 := ListNode{2, nil}
	n1.Next = &ListNode{4, nil}
	n1.Next.Next = &ListNode{3, nil}

	n2 := ListNode{2, nil}
	n2.Next = &ListNode{4, nil}
	n2.Next.Next = &ListNode{3, nil}
	res := Add2NumbersEntry(&n1, &n2)
	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}


func TestRemoveNthFromEndEntry(t *testing.T) {

	n1 := ListNode{1, nil}
	n1.Next = &ListNode{2, nil}
	n1.Next.Next = &ListNode{3, nil}
	n1.Next.Next.Next = &ListNode{4, nil}
	n1.Next.Next.Next.Next = &ListNode{5, nil}
	n1.Next.Next.Next.Next.Next = &ListNode{6, nil}

	res := removeNthFromEnd(&n1, 0)

	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}
