package main

import (
	"fmt"
	"github.com/tcristi74/go-alg/StrStr_28"
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

	fmt.Println("start")

	ret:=StrStr_28.StrStrHash("ends this is the end","end")

	fmt.Println(ret)

}
