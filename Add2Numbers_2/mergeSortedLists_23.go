package Add2Numbers_2

import (
//"flag"
//"math"
//"path/filepath"
)

/**
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/
func mergeKLists(lists []*ListNode) *ListNode {

	iniList := lists[0]
	for i := 1; i < len(lists); i++ {
		iniList = Merge2Lists(iniList, lists[i])
	}
	return iniList
}

func Merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {

	//we need to start with the list with the smaller first element
	var flist, slist *ListNode
	if list1.Val <= list2.Val {
		flist = list1
		slist = list2
	} else {
		flist = list2
		slist = list1
	}

	goOn := true
	var fPointer, sPointer *ListNode
	//var i int = 0
	//var j int = 0
	//moveSecond := false
	for goOn {

		if flist.Val <= slist.Val {
			// we move flist
			// we keep slist the same
			//flist = flist.Next
		} else {
			if flist.Next != nil {

				fNextval := flist.Next.Val
				fPointer = flist.Next
				sPointer = slist
				for true {
					flist.Next = sPointer
					flist = flist.Next

					if slist.Val > fNextval {
						break
					}
					if slist.Next == nil {
						break
					}
					slist = slist.Next
					sPointer = slist
				}
				flist.Next = fPointer
			}

		}

		if flist.Next == nil {
			flist.Next = slist
			goOn = false
		}
		flist = flist.Next

	}
	return flist
}
