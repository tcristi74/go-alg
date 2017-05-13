package Array

import (
	"fmt"
	"testing"
)

//TestGetBestTradingDeal
func TestGetBestTradingDeal(t *testing.T) {

	nums := []int{90, 60, 35, 205, 180, 40, 25, 190}
	fmt.Println(nums)
	ret := GetBestTradingDeal(nums)
	if ret != 170 {
		t.Errorf("BestBet == %d, want %d", ret, 170)
	} else {
		fmt.Printf("Best bet is %d\n", ret)
	}
}

//TestJump45
func TestJump45(t *testing.T) {

	nums := []int{2,3,1,1,4}
	fmt.Println(nums)
	ret := Jump2(nums)
	if ret != 2 {
		t.Errorf("Jump == %d, want %d", ret, 2)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}

	nums = []int{2,3,2,0,1,1,4,2,9}
	fmt.Println(nums)
	ret = Jump2(nums)
	if ret != 5 {
		t.Errorf("Jump == %d, want %d", ret, 5)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}
	//
	nums = []int{2,0,8,0,3,4,7,5,6,1,0,0,5,9,7,5,3,6}
	fmt.Println(nums)
	ret = Jump2(nums)
	if ret != 4 {
		t.Errorf("Jump == %d, want %d", ret, 4)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}


}