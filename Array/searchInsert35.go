package Array

import "sort"

//import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums)==0  || nums[0] > target {
		return 0
	} else 	if nums[len(nums)-1] < target {
		return len(nums)
	}
	return findPosition(&nums,target,0,len(nums)-1, len(nums)/2)
}

func findPosition(nums *[]int, target int, lo int,hi int, index int) int {
	//fmt.Printf("lo=%d	hi=%d	index=%d\n",lo,hi,index)
	if lo+1 == hi {
		if target ==(*nums)[hi] {
			return hi
		} else if target ==(*nums)[lo] {
			return lo
		} else if target > (*nums)[hi] {
			return hi+ 1
		} else if target < (*nums)[lo] {
			return lo - 1
		} else {
			// in between it should be hi
			return hi
		}

	} else {
		idxval := (*nums)[index]
		if target > idxval {
			return findPosition(nums, target, index, hi, (index+hi)/2)
		} else if target< idxval {
			return findPosition(nums, target, lo, index, (lo+index)/2)
		} else {
			return index
		}

	}
}

func searchInsert2(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if i < len(nums) {
		return i
	} else {
		return len(nums)
	}
}



