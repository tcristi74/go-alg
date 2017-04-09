package main

import (
	"fmt"
)
//import "container/list"

var matrix [3][4]string

//var dic *list.List
var dic map[string]bool

func init() {



	// populate matrix

	matrix[0][0]="d"
	matrix[0][1]="d"
	matrix[0][2]="e"
	matrix[0][3]="d"

	matrix[1][0]="e"
	matrix[1][1]="d"
	matrix[1][2]="e"
	matrix[1][3]="d"

	matrix[2][0]="c"
	matrix[2][1]="m"
	matrix[2][2]="r"
	matrix[2][3]="w"


	// populate dictionary
	dic = make(map[string]bool)
	dic["de"] = true
	dic["dec"] = true
	dic["deca"] = true
	dic["decal"] = true
	dic["do"] = true
	dic["dor"] = true
	dic["di"] = true
	dic["dir"] = true

	//dic = list.New()
	//dic.PushBack("de")1//dic.PushBack("dec")
	//dic.PushBack("decal")
	//dic.PushBack("dor")
	//dic.PushBack("di")
	//dic.PushBack("dir")

}

func main() {
	for key,value :=range dic {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
}
