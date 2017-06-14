package stringsAlg

import (
	"sort"
	"sync"
)

type Branch struct {
	sync.RWMutex
	Branches  map[byte]*Branch
	LeafValue []byte
	End       bool
	Count     int64
}

type Trie struct {
	Root *Branch
}

func longestCommonPrefix(words []string) string {

	if len(words)==0 {
		return ""
	}
	if len(words)==1 {
		return words[0]
	}

	sort.Strings(words)
	var maxpre string

	strIni:=""
	for _, str := range words {
		tpre := getMaxPre(strIni, str)
		if len(tpre) > len(maxpre) {
			maxpre = tpre
		}
		strIni = str
	}
	return maxpre
}

func getMaxPre(word1 string, word2 string) string {
	pre := ""
	for i := 0; i < len(word1) && i < len(word2); i++ {
		if string(word1[i]) == string(word2[i]) {
			pre += string(word1[i])
		} else {
			return pre
		}
	}
	return pre
}
