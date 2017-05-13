package Add2Numbers_2



//import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//Given a linked list, remove the nth node from the end of list and return its head.
//For example,
//Given linked list: 1->2->3->4->5, and n = 2.
//After removing the second node from the end, the linked list becomes 1->2->3->5.
func RemoveNthFromEndEntry(head *ListNode, n int) *ListNode {

	ret := removeNthFromEnd(head, n)
	return ret

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n==0 {
		return head
	}

	goon := true
	last := head
	current :=0
	var cn *ListNode
	for goon {
		if current-n== 0{
			cn = head
		} else if current-n > 0 {
			cn = cn.Next
		}
		current++
		if last.Next == nil {
			goon = false
		}
	}
	if cn !=nil && cn.Next != nil  &&  cn.Next.Next != nil {
		cn.Next = cn.Next.Next
	}
	if current==1 && n==1 {
		return new(ListNode)
	}

	return head
}

func Add2NumbersEntry(l1 *ListNode, l2 *ListNode) *ListNode {
	//fmt.Println("l1=",*l1)
	//fmt.Println("l2=",*l2)

	var ret = addTwoNumbers(l1, l2)
	//fmt.Println("result=",*ret)
	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nList := ListNode{}
	last := &nList
	goon := true
	lastIsRoot := true
	var rest int
	for goon {
		if l1 == nil && l2 == nil {
			goon = false
		} else {

			var v1, v2 int
			// calculate value
			if l1 != nil {
				v1 = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				v2 = l2.Val
				l2 = l2.Next
			}
			if !lastIsRoot {

				l := ListNode{}
				last.Next = &l
				//last = l
				last = last.Next
			}

			last.Val, rest = GetVal(v1, v2, rest)
			//add this node to the last node
			lastIsRoot = false
		}
	}
	if rest > 0 {
		l := ListNode{}
		last.Next = &l
		last.Next.Val = rest
	}
	return &nList
}

func GetVal(x, y, r int) (int, int) {
	r1 := (x + y + r) / 10
	m := (x + y + r) % 10
	return m, r1
}
