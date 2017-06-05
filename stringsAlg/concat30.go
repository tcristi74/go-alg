package stringsAlg

//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//For example, given:
//s: "barfoothefoobarman"
//words: ["foo", "bar"]
//
//You should return the indices: [0,9].
//(order does not matter).

type mycounter [2]int

func findSubstring(s string, words []string) []int {
	retVals := make([]int, 0)

	// add words
	ht := make(map[string]mycounter, len(words))
	wLen := len(words[0])

	for i := 0; i < len(words); i++ {

		if val, ok := ht[words[i]]; ok {
			cnt := ht[words[i]]
			cnt[0] = val[0] + 1
			ht[words[i]] = cnt
		} else {
			cnt := ht[words[i]]
			cnt[0] = 1
			ht[words[i]] = cnt
		}
	}
	//fmt.Println("ht=",ht)
	myword := ""
	for i := 0; i <= len(s)-wLen; i++ {
		myword = s[i : i+wLen]
		//check if word is part of words
		_, ok := ht[myword]
		if ok {

			if checkWords(i, myword, s, &ht, len(words)) {
				//fmt.Println("good index",i)
				retVals = append(retVals, i)
			}
		}
	}
	return retVals

}

func checkWords(idx int, word string, s string, ht *map[string]mycounter, wcnt int) bool {

	cleanStatusHt(ht)
	wLen := len(word)
	idx = idx + wLen
	sLen := len(s)
	//decrease one we know it is there
	val, _ := (*ht)[word]
	cnt := (*ht)[word]
	cnt[1] = val[1] + 1
	(*ht)[word] = cnt
	//wcnt := len(*ht) - 1
	for wcnt-1 > 0 {

		if sLen < idx+wLen {
			return false
		}
		nextWord := s[idx : idx+wLen]

		if val, ok := (*ht)[nextWord]; ok {
			//do something here
			if val[1] < val[0] {
				//			fmt.Println("b",*ht)
				// good
				cnt = (*ht)[nextWord]
				cnt[1] = val[1] + 1
				(*ht)[nextWord] = cnt
				idx += wLen
				wcnt--
				//			fmt.Println("e",*ht)
			} else {
				// not good
				return false
			}

		} else {
			// not good
			return false
		}
	}
	//fmt.Println("exit positive fr word:",word)
	return true

}

func cleanStatusHt(ht *map[string]mycounter) {
	for k := range *ht {
		cnt := (*ht)[k]
		cnt[1] = 0
		(*ht)[k] = cnt
	}
}
