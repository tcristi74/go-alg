package Array

import (
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {

	//sort candidates
	sort.Ints(candidates)
	values := [][]int{}

	tempList := make([]int, 0)
	getValidRec(candidates, &values, tempList, 0, target)
	//fmt.Println(values)
	return values
}

func getValidRec(arr []int, result *[][]int, candidate []int, start int, rest int) bool {

	if rest < 0 {
		return false
	} else if rest == 0 {
		//fmt.Println(candidate)
		newc := make([]int, len(candidate))
		copy(newc, candidate)
		//newc = append(newc, candidate...)
		*result = append(*result, newc)
		//fmt.Println(result)
	} else {

		for i := start; i < len(arr); i++ {
			candidate = append(candidate, (arr)[i])
			if !getValidRec(arr, result, candidate, i, rest-(arr)[i]) {
				break
			}
			candidate = (candidate)[:len(candidate)-1]
		}
	}
	return true
}
