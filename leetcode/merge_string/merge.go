package merge_string

func mergeAlternately(word1 string, word2 string) string {
	var resu []uint8
	var length1 = len(word1)
	var length2 = len(word2)
	var bigger = getMax(length1, length2)
	for i := 0; i < bigger; i++ {
		if i < length1 {
			resu = append(resu, word1[i])
		}
		if i < length2 {
			resu = append(resu, word2[i])
		}
	}
	var t = string(resu)
	return t
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
