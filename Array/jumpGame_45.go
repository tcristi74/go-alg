package Array

import (
	"fmt"
)

//Given an array of non-negative integers, you are initially positioned at the first index of the array.
//Each element in the array represents your maximum jump length at that position.
//Your goal is to reach the last index in the minimum number of jumps.
//For example:
//Given array A = [2,3,1,1,4]
//The minimum number of jumps to reach the last index is 2.
//(Jump 1 step from index 0 to 1, then 3 steps to the last index.)
//You can assume that you can always reach the last index.
func Jump(nums []int) int {

	posibHt := make(map[int]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			posibHt[i] = 0
		} else {

			distanceLeft := len(nums) - 1 - i
			dVal := nums[i]
			if dVal >= distanceLeft {
				posibHt[i] = 1
			} else if dVal == 0 {
				posibHt[i] = 99999999
			} else {
				// get best posibility
				minMp := posibHt[i+dVal] + 1

				for ; dVal > 0; dVal-- {

					tval := posibHt[i+dVal]
					if tval < minMp {
						minMp = tval
					}
				}

				posibHt[i] = minMp + 1
			}
		}
	}
	fmt.Println(posibHt)
	return posibHt[0]
}

// from beginning to end
func Jump2(nums []int) int {

	//posibHt := make(map[int]int, len(nums))
	it := 0
	if len(nums) == 1 {
		return it
	}

	rg := [2]int{0, nums[0]}
	fmt.Println(rg)
	min := len(nums)
	max := 0
	for rg[1] < len(nums)-1 && rg[0] < rg[1] {

		//expand the array
		for i := rg[0] + 1; i <= rg[1]; i++ {
			if nums[i] > 0 && i < min {
				min = i
			}
			if max < nums[i]+i {
				max = nums[i] + i
			}
		}
		rg[0] = min
		rg[1] = max

		min = len(nums)
		fmt.Println(rg)
		it++
	}
	return it + 1
}
