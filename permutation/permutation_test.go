package permutation

import (
	"testing"
	"fmt"
)

func TestPermutation(t *testing.T){

	input :=[]int{1,2,3,4}
	ret :=Permute(input)

	fmt.Println(ret)

	if len(ret) != 6 {
		t.Errorf("Length of array must be 6")
	}

}

func TestRestArray(t *testing.T){

	input :=[]int{1,2,3}
	ret :=RestArray(input,2)
	fmt.Println(ret)


}

