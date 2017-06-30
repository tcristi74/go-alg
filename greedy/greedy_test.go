package greedy

import (
	"testing"
	"fmt"
)



func TestJumpGame(t *testing.T){
	arr :=[]int {2,3,1,1,4}
     	ret :=canJump(arr)
	fmt.Println(ret)
}

func TestJumpGame2(t *testing.T){
	arr :=[]int {2,3,1,0,4}
	ret :=canJump(arr)
	fmt.Println(ret)
}
