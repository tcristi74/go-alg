package algorithm

import (
	"fmt"
	"testing"
	"github.com/tcristi74/go-alg/tests"
)


func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr :=tests.LoadArray1()
		_ =MaxSubArray(arr)
	}
}

//TestMaxSubArray  execute with go test -bench=.
func TestMaxSubArray(t *testing.T)  {



	arr :=tests.LoadArray1()
	fmt.Println(arr)
	sum :=MaxSubArray(arr)
	fmt.Println(sum)

	arr=tests.LoadArray2()
	fmt.Println(arr)
	sum = MaxSubArray(arr)
	fmt.Println(sum)

	arr1 := []int {-2,-3,-1}
	fmt.Println(arr1)
	sum = MaxSubArray(arr1)
	fmt.Println(sum)

}
