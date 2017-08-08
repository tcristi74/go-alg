package StrStr_28

func StrStr(haystack string, needle string) int {
	for x,r :=range haystack {
		if string(r)==needle[0:1]{
			xs:=x
			for _,n :=range needle {
				if rune(haystack[xs])!=rune(n) {
					break
				}
				xs++
			}
			if xs==x+len(needle){
				return x
			}
		}
	}
	return -1
}

func StrStrHash(haystack string, needle string) int {
	ht:=make(map[int]int,len(haystack))

	for i :=range haystack{
		if i>len(haystack)-len(needle) {
			break
		}
		hash:=getHash(haystack[i:i+len(needle)])

		if _, ok := ht[hash];!ok {
			ht[hash]=i
		}
	}

	//fmt.Println(ht)
	nHash:=getHash(needle)
	if i, ok := ht[nHash];ok {
		return i
	} 	else {
		return -1
	}
	return ht[getHash(needle)]
}


func getHash(needle string) int{
	//fmt.Println("starthash")
	var ret =0
	for i,r :=range needle {
		ret+=int(r)+i*3
	}

	return ret
}
