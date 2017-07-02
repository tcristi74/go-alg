package main

import (
	"fmt"
	"github.com/tcristi74/go-alg/Hard"
)

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
//func (p *Curves) AddCurve(iniidx int, finidx int, area int) {
//	c := Curve{IniIndex: iniidx, FinIndex: finidx, Area: area}
//	(*p).List = append((*p).List, c)
//	//return p
//}

func main() {
	arr :=[]int {0,1,0,2,1,0,1,3,2,1,2,1}
	//arr =[]int {4,2,3}
	arr =[]int {4,2,3,1,2}
	//arr = []int {4,9,4,5,3,2}
	ret := Hard.Trap(arr)
	fmt.Println(ret)
	if ret !=6 {
		fmt.Println("return must be 6")
	}
}
