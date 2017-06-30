package matrix

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {

	m := make([][]int, 0)
	m = append(m, []int{1, 2, 3})
	m = append(m, []int{2, 5, 6})
	m = append(m, []int{1, 1, 9})
	fmt.Println(m)
	res := minPathSum(m)
	fmt.Println(res)
}
