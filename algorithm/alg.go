package algorithm

import (
	"fmt"
	"github.com/tcristi74/go-alg/tests"
)

var arr []int

//MaxSubArrayEntry
func MaxSubArrayEntry() {
	arr := tests.LoadArray1()
	fmt.Println(arr)
	sum1 := MaxSubArray(arr)
	fmt.Println(sum1)

}

//MaxSubArray
func MaxSubArray(nums []int) int {
	sum := 0
	maxsum := -1000

	curStart := 0

	for i := 0; i < len(nums); i++ {

		if i > curStart && nums[curStart] < 0 {
			fmt.Println(i, curStart, nums[curStart], sum)
			sum = sum - nums[curStart]
			if curStart == 0 {
				maxsum = -1000
			}

			fmt.Println("new sum", sum, maxsum)
			if sum > maxsum  && curStart>0 {
				maxsum = sum
			}
			fmt.Println("new maxsum", maxsum)

			curStart++
		}

		if sum < 0 && nums[i] > 0 {
			curStart = i
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > maxsum {
			maxsum = sum
		}

	}
	return maxsum
}
