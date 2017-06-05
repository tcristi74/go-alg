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

	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(nums)
	ret := Jump2(nums)
	if ret != 2 {
		t.Errorf("Jump == %d, want %d", ret, 2)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}

	nums = []int{2, 3, 2, 0, 1, 1, 4, 2, 9}
	fmt.Println(nums)
	ret = Jump2(nums)
	if ret != 5 {
		t.Errorf("Jump == %d, want %d", ret, 5)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}
	//
	nums = []int{2, 0, 8, 0, 3, 4, 7, 5, 6, 1, 0, 0, 5, 9, 7, 5, 3, 6}
	fmt.Println(nums)
	ret = Jump2(nums)
	if ret != 4 {
		t.Errorf("Jump == %d, want %d", ret, 4)
	} else {
		fmt.Printf("Jump %d\n", ret)
	}

}

//TestRotateArray
func TestRotateArray(t *testing.T) {

	arr := []int{9, 10, 11, 14, 2, 3, 4, 5, 6, 7, 8}
	idx := findMin_153(arr)
	fmt.Println(arr)
	fmt.Printf("index=%d\n", idx)
	if idx != 4 {
		t.Errorf("index =%d, we want %d", idx, 4)
	}
	arr = []int{9, 10, 11, 2, 3, 4, 5, 6, 7, 8}
	idx = findMin_153(arr)
	fmt.Println(arr)
	fmt.Printf("index=%d\n", idx)
	if idx != 3 {
		t.Errorf("index =%d, we want %d", idx, 3)
	}

	arr = []int{9, 10, 11}
	idx = findMin_153(arr)
	fmt.Println(arr)
	fmt.Printf("index=%d\n", idx)
	if idx != 0 {
		t.Errorf("index =%d, we want %d", idx, 0)
	}

	arr = []int{9, 8, 7}
	idx = findMin_153(arr)
	fmt.Println(arr)
	fmt.Printf("index=%d\n", idx)
	if idx != -1 {
		t.Errorf("index =%d, we want %d", idx, -1)
	}

}

func TestCombination39(t *testing.T) {
	nums := []int{2, 3, 6, 7}
	sum := 7
	vals := CombinationSum(nums, sum)

	for i := 0; i < len(vals); i++ {
		fmt.Println(vals[i])
	}
	if len(vals) != 2 {
		t.Errorf("Combination39 Expected results == %d, want %d", 2, 2)
	}
}
func TestCombination39_2(t *testing.T) {
	nums := []int{1, 2, 3, 6, 7}
	sum := 17
	vals := CombinationSum(nums, sum)

	for i := 0; i < len(vals); i++ {
		fmt.Println(vals[i])
	}
	if len(vals) != 75 {
		t.Errorf("Combination39 Expected results == %d, want %d", len(vals), 75)
	}
}

func BenchmarkSum(b *testing.B) {

	nums := []int{1, 2, 3, 6, 7}
	sum := 17
	for i := 0; i < b.N; i++ {
		_ = CombinationSum(nums, sum)
	}
}
