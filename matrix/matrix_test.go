package matrix

import (
	"testing"
	"fmt"
)

func TestSpiralMatrixM1(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1,2})
	m= append(m,[]int {3,4})
	m= append(m,[]int {5,6})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}


func TestSpiralMatrixM2(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1})
	m= append(m,[]int {2})
	m= append(m,[]int {3})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}


func TestSpiralMatrixN1(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1,2,3})
	m= append(m,[]int {4,5,6})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}


func TestSpiralMatrixN2(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1,2,3})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}

func TestSpiralMatrix3(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1,2,3})
	m= append(m,[]int {4,5,6})
	m= append(m,[]int {7,8,9})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}

func TestSpiralMatrix5(t *testing.T){
	m  := make([][]int,0)
	m = append(m,[]int {1,2,3,4,5})
	m= append(m,[]int {6,7,8,9,10})
	m= append(m,[]int {11,12,13,14,15})
	m= append(m,[]int {16,17,18,19,20})
	m= append(m,[]int {21,22,23,24,25})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}


func TestSpiralMatrix2(t *testing.T){


	m  := make([][]int,0)

	m = append(m,[]int {1,2})
	m= append(m,[]int {3,4})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}

func TestSpiralMatrix1(t *testing.T){


	m  := make([][]int,0)

	m = append(m,[]int {1})
	fmt.Println(m)
	ret:=spiralOrder(m)
	fmt.Println(ret)
}