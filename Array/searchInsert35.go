package Array


func searchInsert(nums []int, target int) int {
	if len(nums)==0 {
		return 0
	}
	return findPosition(&nums,target,0,len(nums)-1, len(nums)/2)
}

func findPosition(nums *[]int, target int, lo int,hi int, index int) int {
	if (*nums)[lo] > target {
		return 0
	} else if (*nums)[hi] < target {
		return hi + 1
	}

	if lo+1 == hi {
		if target > (*nums)[index] {
			return index - 1
		} else if target < (*nums)[index] {
			return index + 1
		} else {
			return index
		}

	} else {
		idxval := (*nums)[index]
		if target > idxval {
			return findPosition(nums, target, index+1, hi, (index+hi)/2)
		} else if target< idxval {
			return findPosition(nums, target, lo, index, (lo+index)/2)
		} else {
			return idxval
		}

	}
}



