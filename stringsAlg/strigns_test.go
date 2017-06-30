package stringsAlg

import (
	"fmt"
	"testing"
)

func TestConcat30(t *testing.T) {

	s := "abdemocadbac"
	words := []string{"cad", "bac"}
	idx := findSubstring(s, words)
	fmt.Println(idx)
	if len(idx) != 1 || idx[0] != 6 {
		t.Errorf("Indexes == %d, want %d", idx[0], 6)
	}
}

func TestConcat30_2(t *testing.T) {

	s := "abdemocadbacdemobaccad"
	words := []string{"cad", "bac"}
	idx := findSubstring(s, words)
	fmt.Println(idx)
	if len(idx) != 2 || idx[0] != 6 || idx[1] != 16 {
		fmt.Println(idx)
		t.Errorf("Indexes == %d, want %d", idx[0], 6)
	}
}

func TestConcat30_3(t *testing.T) {

	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	idx := findSubstring(s, words)
	fmt.Println(idx)
	if len(idx) != 1 || idx[0] != 8 {
		fmt.Println(idx)
		t.Errorf("Indexes == %d, want %d", idx[0], 6)
	}
}

func TestLongestValidP(t *testing.T) {
	//s:=")())((()))()))()()"
	s := "(())"
	println(s)
	idx := longestValidParentheses(s)
	expect := 4
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

	s = "(())((("
	println(s)
	idx = longestValidParentheses(s)
	expect = 4
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

	s = ")))"
	println(s)
	idx = longestValidParentheses(s)
	expect = 0
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

	s = "((((()"
	println(s)
	idx = longestValidParentheses(s)
	expect = 2
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

	s = "()()()"
	println(s)
	idx = longestValidParentheses(s)
	expect = 6
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}
	s = "(())()()()"
	println(s)
	idx = longestValidParentheses(s)
	expect = 10
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

	s = "(())()())()"
	println(s)
	idx = longestValidParentheses(s)
	expect = 8
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}

}

func TestLongestValidP2(t *testing.T) {
	s := "(())()())()(((())))"
	println(s)
	idx := longestValidParentheses(s)
	expect := 10
	if idx != expect {
		t.Errorf("Indexes == %d, want %d", idx, expect)
	}
}
