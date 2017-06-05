package Array

import (
	//"fmt"
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {

	if len(s) == 0 && len(p) > 0 || len(p) == 0 && len(s) > 0 {
		return false
	}

	for i, v := range p {
		// check the value
		if len(s) == 0 {
			return false
		}

		switch string(v) {
		case "?":
			fmt.Printf("i=%d, v=%v,  s=%v\n", i, string(v), s)
			s = s[1:]
			//fmt.Println(s)
		case "*":
			fmt.Printf("i=%d, v=%v,  s=%v\n", i, string(v), s)
			// check if is the last string
			if i == len(p)-1 {
				return true
			} else {
				//check after
				if string(p[i+1]) != "*" && string(p[i+1]) != "?" {
					fmt.Printf("looking for: %v \n", string(p[i+1]))
					idxr := strings.IndexRune(s[1:], rune(p[i+1]))
					if idxr >= 0 {
						fmt.Println("got it at ", idxr)
						// we are good
						s = s[idxr+1:]
						fmt.Println(s)
					} else {
						return false
					}
				}
			}

		default:
			fmt.Printf("i=%d, v=%v,  s=%v\n", i, string(v), s)
			if rune(s[0]) != v {
				return false
			} else {
				s = s[1:]
			}
		}
	}

	return len(s) == 0

}
