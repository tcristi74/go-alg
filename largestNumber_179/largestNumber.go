package largestNumber

import (
	"fmt"
	"strconv"
)

type fn func(int, int) int

//Given a list of non negative integers, arrange them such that they form the largest numr.
//For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.
//Note: The result may be very large, so you need to return a string instead of an integer.
func GetLargest(nums []int) string {
	//fmt.Println(nums)
	sortArray(&nums, 0, len(nums)-1)
	//fmt.Println(nums)
	var ret string
	for i := 0; i < len(nums); i++ {
		ret += strconv.FormatInt(int64(nums[i]), 10)
	}
	//fmt.Println(ret)
	return ret
}

func Compare(no1 int, no2 int) int {
	no1s := strconv.FormatInt(int64(no1), 10) +
		strconv.FormatInt(int64(no2), 10)

	no2s := strconv.FormatInt(int64(no2), 10) +
		strconv.FormatInt(int64(no1), 10)

	no1X, err := strconv.ParseInt(no1s, 10, 64)

	if err != nil {
		fmt.Println(err)
		return -1
	}
	no2X, err := strconv.ParseInt(no2s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	if no1X < no2X {
		return -1
	} else if no1X > no2X {
		return 1
	} else {
		return 0
	}
}

func sortArray(nums *[]int, min int, max int) {
	if min < max {
		pivot := partition(nums, min, max, Compare)
		sortArray(nums, min, pivot-1)
		sortArray(nums, pivot+1, max)
	}
}

func partition(nums *[]int, min int, max int, comp fn) int {
	pivot := (*nums)[max]
	i := min - 1
	for j := min; j < max; j++ {
		if comp((*nums)[j], pivot) > 0 {
			i++
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		}
	}
	i++
	(*nums)[i], (*nums)[max] = (*nums)[max], (*nums)[i]
	return i
}
