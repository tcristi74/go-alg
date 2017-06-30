package Array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := ThreeSum(nums)
	fmt.Println(ret)
	if len(ret) != 2 {
		t.Errorf("Length of array must be 2")
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}

func TestThreeSumE1(t *testing.T) {

	nums := []int{0, 0, 0}

	ret := ThreeSum(nums)

	fmt.Println(ret)
	if len(ret) != 1 {
		t.Errorf("Length of array must be 1")
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i])
	}
}
