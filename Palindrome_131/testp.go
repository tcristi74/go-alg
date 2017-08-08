package Palindrome_131

import (
	"testing"
	"fmt"
)

//TestPalin2
func TestPalin2(t *testing.T){

		s:="demo"
		ret:=IsPalin(s)
		if !ret {
			t.Errorf("%s not palindrome",s)
		}
		fmt.Println(ret)
}
