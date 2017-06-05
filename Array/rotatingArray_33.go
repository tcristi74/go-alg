package Array

//import "fmt"

func findMin_153(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return 0
	}
	return getIndex(nums, 0, len(nums)-1, len(nums)/2)
}

func getIndex(nums []int, lo int, hi int, p int) int {
	//fmt.Printf("lo=%d, hi=%d, pivot=%d\n",lo,hi,p)
	if nums[lo] <= nums[hi] {
		return lo
	} else if lo+1 == hi {
		return hi
	} else if nums[lo] < nums[p-1] && nums[p] < nums[hi] {
		return p
	} else if nums[lo] > nums[p] && nums[p] > nums[hi] {
		return -1
	} else {
		//new pivot
		if nums[lo] > nums[p] {
			return getIndex(nums, lo, p, (lo+p)/2)
		} else {
			return getIndex(nums, p, hi, (p+hi)/2)
		}
	}
}
