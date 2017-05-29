package Array

import "testing"

func TestWildcard(t *testing.T) {
	var mystring string
	var pattern string

	//case 1
	//mystring = "abc"
	//pattern = "a??"

	//if !isMatch(mystring, pattern) {
	//
	//	t.Errorf("Wilcard match fails for string== %v, pattern %v", mystring, pattern)
	//}

	//mystring = "abcdeyyyy"
	//pattern = "a???e*"

	mystring="zacabz"
	pattern="*a?b*"

	if !isMatch(mystring, pattern) {

		t.Errorf("Wilcard match fails for string== %v, pattern %v", mystring, pattern)
	}


}

//func TestWildcardFail(t *testing.T) {
//	var mystring string
//	var pattern string
//
//	//case 1
//	mystring = "abc"
//	pattern = "a??"
//
//	if isMatch(mystring, pattern) {
//
//		t.Errorf("Wilcard should NOT match fails for string== %v, pattern %v", mystring, pattern)
//	}
//
//
//}
