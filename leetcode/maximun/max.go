package maximun

import (
	"sort"
	"strings"
)

func maxProduct(words []string) int {
	var max  = 0
	var l []string
	for _, s2 := range words {
		l = append(l,s2)
	}
	sort.Strings(l)

	for i := 0; i <len(l) -1 ; i++ {
		for j := i+1; j <len(l); j++ {
			if !contain(l[j],l[i]){
				var t =  len(l[j]) * len(l[i])
				if t > max{
					max = t
				}
			}
		}
	}
	return max
}

func contain(s1,s2 string) bool{
	for i := 0; i < len(s1); i++ {
		if strings.ContainsRune(s2, rune(s1[i])){
			return true
		}
	}
	return false
}
