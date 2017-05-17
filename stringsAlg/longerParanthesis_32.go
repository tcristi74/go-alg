package stringsAlg

//import (
//	"errors"
//
//)

//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//For "(()", the longest valid parentheses substring is "()", which has length = 2.
//Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.
func longestValidParentheses(s string) int {
	ms := NewStack()
	maxCorrect := 0
	validPoints :=make(map[int]int,0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			//remove from stack, if nothging else reinitialize
			idx, err := ms.Pop()
			if err {
				// clear valid points
				for k := range validPoints {
					delete(validPoints, k)
				}
			} else {

				if  val,ok :=  validPoints[idx-1];ok {
					validPoints[i] = i - idx + 1 +val
				} else {
					validPoints[i] = i - idx + 1
				}
				if validPoints[i] > maxCorrect {
					maxCorrect = validPoints[i]
				}
			}
		} else {
			ms.Push(i)
		}
	}
	return maxCorrect
}
type stack struct {
	s    []int
	Size int
}

func NewStack() *stack {
	v := stack{s: make([]int, 0), Size: 0}
	return &v
}
func (s *stack) Push(idx int) {
	s.s = append(s.s, idx)
	s.Size++
}
func (s *stack) Pop() (int, bool) {
	if s.Size == 0 {
		return 0, true
	}
	res := s.s[s.Size-1]
	s.s = s.s[:s.Size-1]
	s.Size--
	return res, false
}
