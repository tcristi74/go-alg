package SearchRange34

import "testing"

//Test
func TestSearchRange(t *testing.T) {

	nums := []int{5, 7, 7, 8, 8, 10}
	ret := SearchRange(nums, 8)

	if ret[0] != 3 || ret[1] != 4 {
		t.Errorf("values are supposed to be 3 and 4, not %d and %d\n", ret[0], ret[1])
	}
}
