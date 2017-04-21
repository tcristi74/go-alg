package main

import (
	"fmt"
)

import (
	"github.com/tcristi74/go-alg/tests"
	"github.com/tcristi74/go-alg/algorithm"
	"github.com/tcristi74/go-alg/palindrome"
)

var matrix [3][4]string

var dic map[string]bool

type direction string

func init() {



}


func main() {
	v:=palin.CheckStringIsPalindrome("thisiht")
	fmt.Println(v);
	ret :=  palin.GetSubPalindromesFromString("thisihttatbbdbcrabdb")
	for ar,no :=range ret {
		fmt.Println(ar,no)
	}
}

func mainsubarray() {

	// populate matrix
	tests.LoadMatrix(&matrix)

	// populate dictionary
	dic = tests.LoadDictionary()

	arr :=[]int {3,6,-7}

	sum :=algorithm.MaxSubArray(arr)
	fmt.Println(sum)

	for key, value := range dic {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])



	goodWords := make(map[string]bool)
	for i := 0; i < len(matrix); i++ {
		for v := 0; v < len(matrix[i]); v++ {
			search(goodWords, [2]int{i, v}, "right", "")
			search(goodWords, [2]int{i, v}, "left", "")
			search(goodWords, [2]int{i, v}, "up", "")
			search(goodWords, [2]int{i, v}, "down", "")
		}
	}
	//print good words
	fmt.Println("goodWords",goodWords)

}

func search(words map[string]bool, position [2]int, dir direction, prevStr string) {
	//make sure the point exists
	if !pointExists(matrix, position) {
		return
	}
	// add new letter
	prevStr += matrix[position[0]][position[1]]


	// we can optimize this
	if dic[prevStr]  && !words[prevStr]{
		words[prevStr] = true
	}
	search(words, getNewPosition(position, dir), dir, prevStr)


}

func pointExists(matrix [3][4]string, position [2]int) bool {
	//add code here
	if position[0] < 0 || position[0] >= len(matrix) {
		return false
	}
	if position[1] < 0 || position[1] >= len(matrix[0]) {
		return false
	}
	return true
}

func getNewPosition(position [2]int, dir direction) [2]int {

	newPosition := position
	switch dir {
	case "right":
		newPosition[1]++
	case "left":
		newPosition[1]--
	case "up":
		newPosition[0]--
	case "down":
		newPosition[0]++
	}
	return newPosition
}
