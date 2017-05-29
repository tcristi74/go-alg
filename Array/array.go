package Array

//import "fmt"

//import (
//	"fmt"
//)

func GetBestTradingDeal(nums []int) int {
	maxDiff := 0
	var a int = 0
	var aVal int = 0
	for i := 0; i < len(nums); i++ {
		//fmt.Println(i)
		if nums[a] > nums[i] {
			a = i
			aVal = 0
		} else {
			aVal = nums[i] - nums[a]
		}
		if aVal > maxDiff {
			maxDiff = aVal
		}
	}
	return maxDiff
}

