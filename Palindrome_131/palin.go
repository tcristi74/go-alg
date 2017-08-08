package Palindrome_131

func IsPalin(s string) bool {

	if s=="" {
		return false
	}
	r:=len(s)-1
	for i:=1;i<len(s);i++{
		if r<=i {
			return true
		}
		if s[i]!=s[r] {
			return false
		}
		r--;
	}
	return true
}