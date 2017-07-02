package Hard

import (
	"testing"
	"fmt"
)




func TestTrapping(t *testing.T){
	arr :=[]int {0,1,0,2,1,0,1,3,2,1,2,1}
	ret :=Trap(arr)
	fmt.Println(ret)
	if ret !=6 {
		t.Errorf("return must be 6")
	}

}
