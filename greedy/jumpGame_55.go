package greedy

import "fmt"

//Given an array of non-negative integers, you are initially positioned at the first index of the array.
//Each element in the array represents your maximum jump length at that position.
//Determine if you are able to reach the last index
//For example:
//A = [2,3,1,1,4], return true.
//A = [3,2,1,0,4], return false.
func canJump(nums []int) bool {
	if nums[0] == 0 {
		return len(nums) == 1

	}

	for i := len(nums) - 2; i >= 0; i-- {
		// look for zero
		if nums[i] == 0 {
			// we got a problem try solve it
			if k := solve(nums[:i]); k >= 0 {
				fmt.Printf("good k=%d\n", k)
				i = k
			} else {
				return false
			}
		}
	}
	return true
}

func solve(nums []int) int {
	fmt.Println(nums)
	j := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > j {
			fmt.Printf("good val=%d\n", nums[i])
			return i
		} else {
			j++
		}
	}
	return -1
}
