package ClosestPalindrome

import (
	"fmt"
	"strconv"
)

//
//Given an integer n, find the closest integer (not including itself), which is a palindrome.
//
//The 'closest' is defined as absolute difference minimized between two integers.
//
//Example 1:
//Input: "123"
//Output: "121"
//Note:
//The input n is a positive integer represented by string, whose length will not exceed 18.
//If there is a tie, return the smaller one as answer.

func NearestPalindromic(n string) string {

	fmt.Println(n)
	slen := len(n)
	if slen < 2 {
		return "n"
	}

	if slen%2 == 0 {
		return "even"
	} else {
		// this is odd
		// go to the middle of it
		sleft := string(n[:slen/2+1])
		fmt.Printf("sleft=%s\n", sleft)
		sleftNo, _ := strconv.ParseInt(sleft,10,64)
		sleftNo++

		fmt.Printf("middleNo=%d\n", sleftNo)
		//add the rest
		rest := strconv.FormatInt(sleftNo,10)[:slen/2]
		fmt.Printf("upper rest=%s\n", rest)
		upperFinalStr := strconv.FormatInt(sleftNo,10) + reverse(rest)
		fmt.Printf("upperfinal=%s\n", upperFinalStr)
		//lower final
		sleftNo=sleftNo-2
		rest = strconv.FormatInt(sleftNo,10)[:slen/2]
		fmt.Printf("lower rest=%s\n", rest)
		lowerFinalStr := strconv.FormatInt(sleftNo,10) + reverse(rest)
		fmt.Printf("lowerfinal=%s\n", lowerFinalStr)
		lower,_:=strconv.ParseInt(lowerFinalStr,10,64)

		upper,_:=strconv.ParseInt(upperFinalStr,10,64)
		n2,_:=strconv.ParseInt(n,10,64)

		if upper-n2<n2-lower {
			return upperFinalStr
		}
		return lowerFinalStr
	}
}

	func reverse( ss string) string{

			newstr:=""
			for i:=len(ss)-1;i>=0;i-- {

				newstr+=string(ss[i])
			}
			return newstr

	}

