package Palindrome_131

import (
	"testing"
	"fmt"
)

//TestCheckStringIsPalindrome
func TestPalin(t *testing.T){
s:="demo"
	ret:=IsPalin(s)
	if !ret {
		t.Errorf("%s not palindrome",s)
	}
	fmt.Println(ret)
}


