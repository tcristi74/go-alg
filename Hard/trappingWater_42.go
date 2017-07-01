package Water

import (
	//"math"
	//"golang.org/x/tools/go/gcimporter15/testdata"
)
import "fmt"

import "math"

//import "strconv"

//Given n non-negative integers representing an elevation map where the width of each bar is 1,

//For example,
//Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.

//type Curve struct {
//	IniIndex int
//	FinIndex int
//	Area     int
//}
//
//type Curves struct {
//	List []Curve
//}
//
//func (p *Curves) AddCurve(iniidx int, finidx int, area int)  {
//	c := Curve{IniIndex: iniidx, FinIndex: finidx, Area: area}
//	(*p).List = append((*p).List, c)
//
//}
//					curvs.AddCurve(s, i, curCurve)

func Trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	ret := 0
	for s := 0; s < len(height); s++ {
		if height[s] == 0 {
			continue
		}
		iniIdx := s

		curCurve := 0
		curveHighSoFar:=0
		curveOn := false
		curveHighSoFar =0
		for i := s+1; i < len(height); i++ {
			if curveOn {
				if height[i] >= height[iniIdx] {
					//curve ends
					curveOn = false
					// add to the final curve
					ret += curCurve
					//set another bar
					s = i
					iniIdx =i
					curCurve=0
					curveHighSoFar =0
				} else {
					curveHighSoFar=int(math.Max(float64(curveHighSoFar), float64(height[i])))
					if len(height)-1==i {
						// this is the end and it is not closed
						s = iniIdx-1
						height[iniIdx] =curveHighSoFar
						fmt.Println(curveHighSoFar)
						//curCurve=0
					} else {
						// adding to
						curCurve += height[iniIdx] - height[i]
					}
				}
			} else {
				if height[iniIdx] > height[i] {
					//start curve
					curveOn = true
					curCurve += height[iniIdx] - height[i]
					curveHighSoFar=int(math.Max(float64(curveHighSoFar), float64(height[i])))
				} else {
					iniIdx = i
					curveOn = false
					curveHighSoFar=0
				}
			}
		}
	}
	return ret
}
