package Array

import (
	"testing"
	"fmt"
)

func TestFindkLargest(t *testing.T) {
	nums :=[]int {3,6,5,2,7,9,10,34,67,1}
	ret:=findKthLargest(nums,4)
	if ret!=9 {
		t.Errorf("bad")
	}
	fmt.Println(ret)
}
