package palin

import (
	"fmt"
)

//var m map[string]int

func GetSubPalindromesFromString(str string) map[string]int {

	m := make(map[string]int)

	var w string
	k := 0
	for i := 1; i <= len(str); i++ {
		w = str[k:i]

		if CheckStringIsPalindrome(w) {
			_, ok := m[w]
			if ok {
				m[w]++
			} else {
				m[w] = 1
			}
		}
		if len(w) > (len(str)+1-k)/2 {
			k++
			i = k
		}
	}
	return m
}

func CheckStringIsPalindrome(str string) bool {

	if len(str) ==01 {
		return false
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	fmt.Println("str is palindrom", str)
	return true

}
