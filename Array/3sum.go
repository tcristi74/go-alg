package Array

import "fmt"

//Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
func ThreeSum(nums []int) [][]int {

	var res = [][]int{}

	if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
		res = append(res, []int{0, 0, 0})
		return res
	}

	neg := make(map[int]bool)
	pos := make(map[int]bool)

	// sort the array in negative and positive
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if _, ok := neg[nums[i]]; !ok {
				neg[nums[i]] = true
			}
		} else {
			if _, ok := pos[nums[i]]; !ok {
				pos[nums[i]] = true
			}
		}
	}

	fmt.Println(neg)
	fmt.Println(pos)

	// get 2 positive equal one neg
	for i, _ := range neg {
		get2sumProblem(nums, -i, &res, false)
	}

	for i, _ := range pos {
		get2sumProblem(nums, -i, &res, true)
	}

	fmt.Println(res)
	fmt.Printf("done")
	return res

}

func get2sumProblem(arr []int, s int, res *[][]int, selN bool) {

	fmt.Printf("get2sumProblem with s=%d and type =%b\n", s, selN)
	fmt.Println(arr)

	dic := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if selN && arr[i] < 0 || !selN && arr[i] >= 0 {

			if idx, ok := dic[arr[i]]; ok {
				a := []int{arr[i], arr[idx], -s}
				fmt.Printf("append ")
				*res = append(*res, a)
			} else {
				dic[s-arr[i]] = i
			}
		}
	}
	fmt.Println(dic)

}
