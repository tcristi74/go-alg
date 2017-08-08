package ClosestPalindrome

import "testing"


func TestClosest(t *testing.T) {

	res := NearestPalindromic("123")

	if res != "121" {
		t.Errorf("%f must be %f\n", res, "121")
	}
}

func TestReverse(t *testing.T) {

	res := reverse("DEMO")

	if res != "OMED" {
		t.Errorf("%f must be %f\n", res, "OMED")
	}

}
