package StrStr_28

import (
	"testing"
	"fmt"
)

func TestNeedle(t *testing.T){

	ret:=StrStr("my only frinds , the etnd ,this is the end","end")
	fmt.Println(ret)
}
func TestNeedleHash(t *testing.T){

	ret:=StrStrHash("my only frinds , the etnd ,this is the end","end")
	fmt.Println(ret)
}