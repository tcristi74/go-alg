package SearchRange34

import "path/filepath"

func SearchRange(nums []int, target int) []int {
	ret :=[]int {-1,-1}
	if len(nums)==0 || len(nums)==1 && nums[0]!=target{
		return ret
	}
	idx :=getIndex(&nums,0,len(nums)-1,len(nums)/2,target)
	if idx==-1 {
		return ret
	}
	// look for limits
	maxidx:=idx
	minidx:=idx

	for idx>0 {
		idx =getIndex(&nums,0,idx-1,idx/2,target)
		if idx>=0 {
			minidx=idx
		}
	}



	return ret
}


func getIndex(nums *[]int,loIdx int, hiIdx int, pidx int,target int) int {
	if (*nums)[pidx] == target {
		return pidx
	}
	if loIdx+1 == hiIdx && (*nums)[loIdx] != target {
		return -1
	} else if loIdx+2 == hiIdx && (*nums)[loIdx] != target && (*nums)[hiIdx] != target {
		return -1
	} else if (*nums)[loIdx] > target {
		return -1
	} else if (*nums)[hiIdx] < target {
		return -1
	} else {
		// further investigation
		if pidx > (*nums)[pidx] {
			return getIndex(nums, pidx, hiIdx, (hiIdx-loIdx)/2, target)
		} else {
			return getIndex(nums, loIdx,pidx, (pidx-loIdx)/2, target)
		}
	}
}

