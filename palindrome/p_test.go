package palin

import (
	"fmt"
	"testing"
)

//TestCheckStringIsPalindrome
func TestCheckStringIsPalindrome(t *testing.T) {

	ret := CheckStringIsPalindrome("aathisasihtaa")
	fmt.Println(ret)
}

func TestGetSubPalindromesFromString(t *testing.T) {
	GetSubPalindromesFromString("thissihmaamdoredoreer")
}
