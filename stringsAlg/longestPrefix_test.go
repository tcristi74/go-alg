package stringsAlg

import (
	"testing"
	"fmt"
)

func TestLongestPrefix(t *testing.T) {

	words :=[]string{"florida","macaro","macaroane","flory", "cristi"}

	pref:= longestCommonPrefix(words)

	fmt.Println(pref)
}
