package max_length_charaters

func MaxLengthBetweenEqualCharacters(s string) int {
	if len(s) == 1 {
		return -1
	}
	maxResult := -1
	var tempMap = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if index, ok := tempMap[s[i]]; !ok {
			tempMap[s[i]] = i
		} else {
			maxResult = getMax(maxResult, i-index-1)
		}
	}
	return maxResult
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
