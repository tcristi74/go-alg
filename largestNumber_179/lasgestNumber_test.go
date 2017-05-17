package largestNumber

import (
	"fmt"
	"testing"
)

//TestGetLargest
func TestGetLargest(t *testing.T) {
	nums := []int{5, 98, 99, 9, 9457, 34, 2}
	res := GetLargest(nums)
	fmt.Println(res)

}

func _TestCompare(t *testing.T) {
	res := Compare(2344, 67)
	fmt.Println(res)

	res = Compare(844, 67)
	fmt.Println(res)

	res = Compare(999, 99)
	fmt.Println(res)

	res = Compare(98, 997)
	fmt.Println(res)

	res = Compare(888, 9)
	fmt.Println(res)

}
