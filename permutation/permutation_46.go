package permutation

import (
	"bytes"
	"strconv"
)

func Permute(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) == 1 {
		ret = append(ret, nums)
		return ret
	}
	ht := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		lead := []int{nums[i]}
		r := RestArray(nums, i)
		perm(r, lead, &ret, &ht)
	}
	return ret
}

func perm(rest []int, lead []int, ret *[][]int, ht *map[string]bool) {
	for i := 0; i < len(rest); i++ {
		lead1 := append(lead, rest[i])
		r := RestArray(rest, i)
		if len(rest) == 1 {
			lead1 = append(lead1, r...)
			lead1str := arrtoString(lead1)
			if _, ok := (*ht)[lead1str]; !ok {
				*ret = append(*ret, lead1)
				(*ht)[lead1str] = true
			}

		} else {
			perm(r, lead1, ret, ht)
		}
	}
}

func RestArray(arr []int, idx int) []int {

	if idx == 0 {
		return (arr)[1:]
	} else if idx == len(arr)-1 {
		return arr[0:idx]
	} else {
		var v []int
		v = append(v, arr[0:idx]...)
		v = append(v, arr[idx+1:]...)
		return v
	}
}

func arrtoString(arr []int) string {
	var buffer bytes.Buffer
	for x:=0;x<len(arr);x++ {
		buffer.WriteString(strconv.Itoa(arr[x]))
		buffer.WriteString("|")
	}
	return buffer.String()
}
