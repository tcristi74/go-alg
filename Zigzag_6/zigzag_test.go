package Zigzag_6

import (
	"fmt"
	"testing"
)

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//P   A   H   N
//A P L S I I G
//Y   I   R
//
//And then read line by line: "PAHNAPLSIIGYIR"

//PAY
//PAL
//ISH
//IRI
//NG


func TestZigzag(t *testing.T) {
	ret := Convert("PAYPALISHIRING", 3)

	fmt.Println(ret)

}
